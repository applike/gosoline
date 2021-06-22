package log

import (
	"fmt"
	"github.com/applike/gosoline/pkg/encoding/json"
	"github.com/fatih/color"
	"strings"
)

type Formatter func(timestamp string, level int, format string, args []interface{}, err error, data Data) ([]byte, error)

var formatters = map[string]Formatter{
	"console": FormatterConsole,
	"json":    FormatterJson,
}

func FormatterConsole(timestamp string, level int, format string, args []interface{}, err error, data Data) ([]byte, error) {
	fieldString := getFieldsAsString(data.Fields)
	contextString := getFieldsAsString(data.ContextFields)

	levelStr := fmt.Sprintf("%-7v", LevelName(level))
	channel := fmt.Sprintf("%-7s", data.Channel)
	msg := fmt.Sprintf(format, args...)

	if err != nil {
		msg = color.RedString(err.Error())
	}

	output := fmt.Sprintf("%s %s %s %-50s %s %s",
		color.YellowString(timestamp),
		color.GreenString(channel),
		color.GreenString(levelStr),
		msg,
		color.GreenString(contextString),
		color.BlueString(fieldString),
	)

	output = strings.TrimSpace(output)
	serialized := []byte(output)

	return append(serialized, '\n'), nil
}

func FormatterJson(timestamp string, level int, format string, args []interface{}, err error, data Data) ([]byte, error) {
	msg := fmt.Sprintf(format, args...)
	jsn := make(map[string]interface{}, 8)

	if err != nil {
		jsn["err"] = err.Error()
	}

	jsn["channel"] = data.Channel
	jsn["level"] = LevelName(level)
	jsn["level_name"] = level
	jsn["timestamp"] = timestamp
	jsn["message"] = msg
	jsn["fields"] = data.Fields
	jsn["context"] = data.ContextFields

	serialized, err := json.Marshal(jsn)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}

	return append(serialized, '\n'), nil
}

func getFieldsAsString(fields map[string]interface{}) string {
	fieldParts := make([]string, 0, len(fields))

	for k, v := range fields {
		fieldParts = append(fieldParts, fmt.Sprintf("%v: %v", k, v))
	}

	return strings.Join(fieldParts, ", ")
}
