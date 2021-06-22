package cli

import (
	"github.com/applike/gosoline/pkg/clock"
	"github.com/applike/gosoline/pkg/log"
	"io"
)

func newCliLogger() (log.Logger, error) {
	var err error
	var writer io.Writer

	if writer, err = log.NewIoWriterFile("logs.log"); err != nil {

	}

	handler := log.NewHandlerIoWriter(log.LevelInfo, []string{}, log.FormatterConsole, "", writer)
	logger := log.NewLoggerWithInterfaces(clock.Provider, []log.Handler{handler})

	return logger, nil
}
