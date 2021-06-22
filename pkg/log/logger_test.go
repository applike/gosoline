package log_test

import (
	"bytes"
	"github.com/applike/gosoline/pkg/clock"
	"github.com/applike/gosoline/pkg/log"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestLoggerIoWriter(t *testing.T) {
	var buf = &bytes.Buffer{}
	var handler = log.NewHandlerIoWriter(log.LevelInfo, []string{"main"}, log.FormatterJson, time.RFC3339, buf)
	var cl = clock.NewFakeClock()

	logger := log.NewLoggerWithInterfaces(cl, []log.Handler{handler})

	logger.Info("foo")
	cl.Advance(time.Minute)
	logger.Info("bar")
	logger.Debug("some debug")
	logger.WithChannel("other channel").Info("something in another channel")
	cl.Advance(time.Minute)
	logger.Info("foobaz")

	lines := getLogLines(buf)
	assert.Len(t, lines, 3)

	assert.JSONEq(t, `{"channel":"main","context":{},"fields":{},"level":"info","level_name":2,"message":"foo","timestamp":"1984-04-04T00:00:00Z"}`, lines[0])
	assert.JSONEq(t, `{"channel":"main","context":{},"fields":{},"level":"info","level_name":2,"message":"bar","timestamp":"1984-04-04T00:01:00Z"}`, lines[1])
	assert.JSONEq(t, `{"channel":"main","context":{},"fields":{},"level":"info","level_name":2,"message":"foobaz","timestamp":"1984-04-04T00:02:00Z"}`, lines[2])
}

func getLogLines(buf *bytes.Buffer) []string {
	lines := make([]string, 0)

	for _, line := range strings.Split(buf.String(), "\n") {
		if len(line) == 0 {
			continue
		}

		lines = append(lines, line)
	}

	return lines
}
