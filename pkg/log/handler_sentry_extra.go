package log

import "github.com/pkg/errors"

type ConfigProvider interface {
	AllSettings() map[string]interface{}
}

type SentryExtraProvider func(config ConfigProvider, sentryHook *HandlerSentry) (*HandlerSentry, error)

func SentryExtraConfigProvider(config ConfigProvider, handler *HandlerSentry) (*HandlerSentry, error) {
	configValues := config.AllSettings()
	handler = handler.WithExtra(map[string]interface{}{
		"config": configValues,
	})

	return handler, nil
}

func SentryExtraEcsMetadataProvider(_ ConfigProvider, handler *HandlerSentry) (*HandlerSentry, error) {
	ecsMetadata, err := ReadEcsMetadata()

	if err != nil {
		return handler, errors.Wrap(err, "can not read ecs metadata")
	}

	if ecsMetadata != nil {
		return handler, nil
	}

	handler = handler.WithExtra(map[string]interface{}{
		"ecsMetadata": ecsMetadata,
	})

	return handler, nil
}
