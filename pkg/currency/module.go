package currency

import (
	"context"
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/kernel"
	"github.com/applike/gosoline/pkg/mon"
	"time"
)

type Module struct {
	kernel.BackgroundModule
	kernel.ServiceStage
	updaterService UpdaterService
	logger         mon.Logger
}

func NewCurrencyModule() kernel.ModuleFactory {
	return func(ctx context.Context, config cfg.Config, logger mon.Logger) (kernel.Module, error) {
		module := &Module{
			logger:         logger,
			updaterService: NewUpdater(config, logger),
		}

		return module, nil
	}
}

func (module *Module) Run(ctx context.Context) error {
	ticker := time.NewTicker(time.Duration(1) * time.Hour)
	module.refresh(ctx)
	for {
		select {
		case <-ctx.Done():
			return nil

		case <-ticker.C:
			module.refresh(ctx)
		}
	}
}

func (module *Module) refresh(ctx context.Context) {
	err := module.updaterService.EnsureRecentExchangeRates(ctx)
	if err != nil {
		module.logger.Error(err, "failed to refresh currency exchange rates")
	}
}
