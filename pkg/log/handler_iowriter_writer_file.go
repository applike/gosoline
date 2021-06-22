package log

import (
	"fmt"
	"github.com/applike/gosoline/pkg/cfg"
	"io"
	"os"
)

type ioWriterFileSettings struct {
	Path string `cfg:"path" default:"logs.log"`
}

func ioWriterFileFactory(config cfg.Config, handlerIndex int) (io.Writer, error) {
	key := fmt.Sprintf("log.handlers[%d]", handlerIndex)
	settings := &ioWriterFileSettings{}
	config.UnmarshalKey(key, settings)

	return NewIoWriterFile(settings.Path)
}

func NewIoWriterFile(path string) (io.Writer, error) {
	var err error
	var file *os.File

	if file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600); err != nil {
		return nil, fmt.Errorf("can not open file %s to write logs to: %w", path, err)
	}

	return file, nil
}
