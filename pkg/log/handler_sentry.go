package log

import (
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/getsentry/sentry-go"
	"time"
)

func init() {
	AddHandlerFactory("sentry", handlerSentryFactory)
}

//go:generate mockery -name Sentry
type Sentry interface {
	CaptureException(exception error, hint *sentry.EventHint, scope sentry.EventModifier) *sentry.EventID
}

func handlerSentryFactory(config cfg.Config, _ int) (Handler, error) {
	return NewHandlerSentry(config), nil
}

type HandlerSentry struct {
	sentry Sentry
	tags   map[string]string
	extra  map[string]interface{}
}

func NewHandlerSentry(config cfg.Config) *HandlerSentry {
	env := config.GetString("env")
	appName := config.GetString("app_name")

	tags := map[string]string{
		"application": appName,
	}

	client, _ := sentry.NewClient(sentry.ClientOptions{
		Environment: env,
	})

	return &HandlerSentry{
		sentry: client,
		tags:   tags,
		extra:  make(map[string]interface{}),
	}
}

func (h HandlerSentry) WithExtra(extra map[string]interface{}) *HandlerSentry {
	newExtra := mergeFields(h.extra, extra)

	return &HandlerSentry{
		sentry: h.sentry,
		extra:  newExtra,
	}
}

func (h HandlerSentry) Channels() []string {
	return []string{}
}

func (h HandlerSentry) Level() int {
	return PriorityError
}

func (h HandlerSentry) Log(_ time.Time, _ int, _ string, _ []interface{}, err error, data Data) error {
	if err == nil {
		return nil
	}

	extra := mergeFields(h.extra, data.Fields)
	extra = mergeFields(extra, data.ContextFields)

	scope := sentry.NewScope()
	scope.SetTags(h.tags)
	scope.SetExtras(extra)

	data.Fields["sentry_event_id"] = h.sentry.CaptureException(err, nil, scope)

	return err
}
