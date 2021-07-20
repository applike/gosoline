package log

import (
	"github.com/applike/gosoline/pkg/clock"
	"os"
)

func NewCliLogger() Logger {
	handler := NewHandlerIoWriter(LevelInfo, []string{}, FormatterConsole, "15:04:05.000", os.Stdout)

	return NewLoggerWithInterfaces(clock.Provider, []Handler{handler})
}
