package mon

import (
	"context"
	"fmt"
	"github.com/jonboulle/clockwork"
	"sync"
	"time"
)

type SamplingLogger struct {
	Logger
	clock    clockwork.Clock
	logs     *sync.Map
	interval time.Duration
}

func NewSamplingLogger(logger Logger, interval time.Duration) Logger {
	clock := clockwork.NewRealClock()

	return NewSamplingLoggerWithInterfaces(logger, clock, interval)
}

func NewSamplingLoggerWithInterfaces(logger Logger, clock clockwork.Clock, interval time.Duration) *SamplingLogger {
	return &SamplingLogger{
		Logger:   logger,
		clock:    clock,
		logs:     &sync.Map{},
		interval: interval,
	}
}

func (l *SamplingLogger) copy(logger Logger) *SamplingLogger {
	return &SamplingLogger{
		Logger:   logger,
		clock:    l.clock,
		logs:     l.logs,
		interval: l.interval,
	}
}

func (l *SamplingLogger) WithChannel(channel string) Logger {
	logger := l.Logger.WithChannel(channel)
	return l.copy(logger)
}

func (l *SamplingLogger) WithContext(ctx context.Context) Logger {
	logger := l.Logger.WithContext(ctx)
	return l.copy(logger)
}

func (l *SamplingLogger) WithFields(fields Fields) Logger {
	logger := l.Logger.WithFields(fields)
	return l.copy(logger)
}

func (l *SamplingLogger) Debug(args ...interface{}) {
	if !l.shouldLog(fmt.Sprint(args...)) {
		return
	}

	l.Logger.Debug(args...)
}

func (l *SamplingLogger) Debugf(msg string, args ...interface{}) {
	if !l.shouldLog(msg) {
		return
	}

	l.Logger.Debugf(msg, args...)
}

func (l *SamplingLogger) Error(err error, msg string) {
	if !l.shouldLog(msg) {
		return
	}

	l.Logger.Error(err, msg)
}

func (l *SamplingLogger) Errorf(err error, msg string, args ...interface{}) {
	if !l.shouldLog(msg) {
		return
	}

	l.Logger.Errorf(err, msg, args...)
}

func (l *SamplingLogger) Info(args ...interface{}) {
	if !l.shouldLog(fmt.Sprint(args...)) {
		return
	}

	l.Logger.Info(args...)
}

func (l *SamplingLogger) Infof(msg string, args ...interface{}) {
	if !l.shouldLog(msg) {
		return
	}

	l.Logger.Infof(msg, args...)
}

func (l *SamplingLogger) Warn(args ...interface{}) {
	if !l.shouldLog(fmt.Sprint(args...)) {
		return
	}

	l.Logger.Warn(args...)
}

func (l *SamplingLogger) Warnf(msg string, args ...interface{}) {
	if !l.shouldLog(msg) {
		return
	}

	l.Logger.Warnf(msg, args...)
}

func (l *SamplingLogger) shouldLog(msg string) bool {
	value, ok := l.logs.Load(msg)

	if !ok {
		lastLoggedAt := l.clock.Now()
		l.logs.Store(msg, &lastLoggedAt)

		return true
	}

	lastLoggedAt := value.(*time.Time)
	shouldLogAgainAt := lastLoggedAt.Add(l.interval)

	if shouldLogAgainAt.UnixNano() >= l.clock.Now().UnixNano() {
		return false
	}

	newLastLoggedAt := l.clock.Now()
	l.logs.Store(msg, &newLastLoggedAt)

	return true
}
