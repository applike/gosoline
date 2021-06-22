package log

import (
	"fmt"
	"github.com/applike/gosoline/pkg/cfg"
	"time"
)

type Handler interface {
	Channels() []string
	Level() int
	Log(timestamp time.Time, level int, msg string, args []interface{}, err error, data Data) error
}

type HandlerFactory func(config cfg.Config, handlerIndex int) (Handler, error)

var handlerFactories = map[string]HandlerFactory{
	"iowriter": handlerIoWriterFactory,
}

func AddHandlerFactory(name string, factory HandlerFactory) {
	handlerFactories[name] = factory
}

func NewHandlersFromConfig(config cfg.Config) ([]Handler, error) {
	settings := &LoggerSettings{}
	config.UnmarshalKey("log", settings)

	var ok bool
	var err error
	var handlerFactory HandlerFactory
	var handlers = make([]Handler, len(settings.Handlers))

	for i, handlerSettings := range settings.Handlers {
		if handlerFactory, ok = handlerFactories[handlerSettings.Type]; !ok {
			return nil, fmt.Errorf("there is no logging handler of type %s", handlerSettings.Type)
		}

		if handlers[i], err = handlerFactory(config, i); err != nil {
			return nil, fmt.Errorf("can not create logging handler of type %s on index %d", handlerSettings.Type, i)
		}
	}

	return handlers, nil
}
