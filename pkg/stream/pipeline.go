package stream

import (
	"context"
	"fmt"
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/coffin"
	"github.com/applike/gosoline/pkg/kernel"
	"github.com/applike/gosoline/pkg/mon"
	"sync"
	"time"
)

const MetricNamePipelineReceivedCount = "PipelineReceivedCount"
const MetricNamePipelineProcessedCount = "PipelineProcessedCount"

//go:generate mockery -name PipelineCallback
type PipelineCallback interface {
	Boot(config cfg.Config, logger mon.Logger) error
	Process(ctx context.Context, messages []*Message) ([]*Message, error)
}

type PipelineSettings struct {
	Interval  time.Duration
	BatchSize int
}

type Pipeline struct {
	kernel.EssentialModule
	kernel.ApplicationStage
	ConsumerAcknowledge

	metric   mon.MetricWriter
	cfn      coffin.Coffin
	lck      sync.Mutex
	encoder  MessageEncoder
	output   Output
	ticker   *time.Ticker
	batch    []*Message
	stages   []PipelineCallback
	settings *PipelineSettings
}

func NewPipeline(stages ...PipelineCallback) *Pipeline {
	return &Pipeline{
		cfn:    coffin.New(),
		stages: stages,
	}
}

func (p *Pipeline) Boot(config cfg.Config, logger mon.Logger) error {
	for _, stage := range p.stages {
		err := stage.Boot(config, logger)

		if err != nil {
			return err
		}
	}

	defaults := getDefaultPipelineMetrics()
	metric := mon.NewMetricDaemonWriter(defaults...)

	input := NewConfigurableInput(config, logger, "pipeline")
	output := NewConfigurableOutput(config, logger, "pipeline")

	settings := &PipelineSettings{
		Interval:  config.GetDuration("pipeline_interval") * time.Second,
		BatchSize: config.GetInt("pipeline_batch_size"),
	}

	return p.BootWithInterfaces(logger, metric, input, output, settings)
}

func (p *Pipeline) BootWithInterfaces(logger mon.Logger, metric mon.MetricWriter, input Input, output Output, settings *PipelineSettings) error {
	p.logger = logger
	p.metric = metric
	p.input = input
	p.output = output
	p.encoder = NewMessageEncoder(&MessageEncoderSettings{})
	p.ticker = time.NewTicker(settings.Interval)
	p.batch = make([]*Message, 0, settings.BatchSize)
	p.settings = settings

	return nil
}

func (p *Pipeline) Run(ctx context.Context) error {
	defer p.logger.Info("leaving pipeline")
	defer p.process(ctx, true)

	p.cfn.GoWithContextf(ctx, p.read, "panic during consuming")
	p.cfn.GoWithContextf(ctx, p.input.Run, "panic during run of the consumer input")

	for {
		select {
		case <-ctx.Done():
			p.input.Stop()
			return p.cfn.Wait()

		case <-p.cfn.Dead():
			p.input.Stop()
			return p.cfn.Err()
		}
	}
}

func (p *Pipeline) read(ctx context.Context) error {
	for {
		force := false

		select {
		case msg, ok := <-p.input.Data():
			if !ok {
				return nil
			}

			disaggregated, err := p.disaggregateMessage(ctx, msg)

			if err != nil {
				p.logger.Errorf(err, "can not disaggregate the message")
				continue
			}

			p.batch = append(p.batch, disaggregated...)
			p.Acknowledge(ctx, msg)

		case <-p.ticker.C:
			force = true
		}

		if len(p.batch) >= p.settings.BatchSize || force {
			p.process(ctx, force)
		}
	}
}

func (p *Pipeline) disaggregateMessage(ctx context.Context, msg *Message) ([]*Message, error) {
	if _, ok := msg.Attributes[AttributeAggregate]; !ok {
		return []*Message{msg}, nil
	}

	batch := make([]*Message, 0)
	_, _, err := p.encoder.Decode(ctx, msg, &batch)

	if err != nil {
		return nil, fmt.Errorf("can not decode message: %w", err)
	}

	return batch, nil
}

func (p *Pipeline) process(ctx context.Context, force bool) {
	p.lck.Lock()
	defer p.lck.Unlock()

	batchSize := len(p.batch)

	if batchSize == 0 {
		p.logger.Info("pipeline has nothing to do")
		return
	}

	if batchSize < p.settings.BatchSize && !force {
		return
	}

	p.metric.WriteOne(&mon.MetricDatum{
		MetricName: MetricNamePipelineReceivedCount,
		Value:      float64(batchSize),
	})

	defer func() {
		p.ticker = time.NewTicker(p.settings.Interval)
		p.batch = make([]*Message, 0, p.settings.BatchSize)

		err := coffin.ResolveRecovery(recover())

		if err == nil {
			return
		}

		p.logger.Error(err, "panic when processing batch")
	}()

	p.ticker.Stop()

	var err error

	for _, stage := range p.stages {
		p.batch, err = stage.Process(ctx, p.batch)

		if err != nil {
			p.logger.Error(err, "could not process the batch")
		}
	}

	err = p.output.Write(ctx, MessagesToWritableMessages(p.batch))

	if err != nil {
		p.logger.Error(err, "could not write messages to output")

		return
	}

	processedCount := len(p.batch)

	p.logger.Infof("pipeline processed %d of %d messages", processedCount, batchSize)
	p.metric.WriteOne(&mon.MetricDatum{
		MetricName: MetricNamePipelineProcessedCount,
		Value:      float64(processedCount),
	})
}

func getDefaultPipelineMetrics() mon.MetricData {
	return mon.MetricData{
		{
			Priority:   mon.PriorityHigh,
			MetricName: MetricNamePipelineReceivedCount,
			Unit:       mon.UnitCount,
			Value:      0.0,
		}, {
			Priority:   mon.PriorityHigh,
			MetricName: MetricNamePipelineProcessedCount,
			Unit:       mon.UnitCount,
			Value:      0.0,
		},
	}
}
