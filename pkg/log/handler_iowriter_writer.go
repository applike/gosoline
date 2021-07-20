package log

import (
	"github.com/applike/gosoline/pkg/cfg"
	"io"
	"os"
)

type IoWriterWriterFactory func(config cfg.Config, configKey string) (io.Writer, error)

var ioWriterFactories = map[string]IoWriterWriterFactory{
	"file":   ioWriterFileFactory,
	"stdout": ioWriterStdOutFactory,
}

func ioWriterStdOutFactory(_ cfg.Config, _ string) (io.Writer, error) {
	return os.Stdout, nil
}
