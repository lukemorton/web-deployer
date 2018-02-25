package log

import (
	"bytes"
	"fmt"
	"sort"
	"sync"

  "github.com/sirupsen/logrus"
)

const (
	nocolor  = "0"
	red      = "31"
	green    = "32"
	yellow   = "33"
	blue     = "36"
  white    = "0;37"
	gray     = "2;37"
)

 // LogrusTextFormatter formats logs into text
type LogrusTextFormatter struct {
	// The fields are sorted by default for a consistent output. For applications
	// that log extremely frequently and don't use the JSON formatter this may not
	// be desired.
	DisableSorting bool

	// QuoteEmptyFields will wrap empty fields in quotes if true
	QuoteEmptyFields bool

	sync.Once
}

// Format renders a single log entry
func (f *LogrusTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	keys := make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		keys = append(keys, k)
	}

	if !f.DisableSorting {
		sort.Strings(keys)
	}
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

  f.printColored(b, entry, keys)

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *LogrusTextFormatter) printColored(b *bytes.Buffer, entry *logrus.Entry, keys []string) {
	var levelColor string
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = white
	}

  fmt.Fprintf(b, "\x1b[%sm%-44s\x1b[0m ", levelColor, entry.Message)

	for _, k := range keys {
		v := entry.Data[k]
		fmt.Fprintf(b, " \x1b[%sm%s\x1b[0m=", levelColor, k)
		f.appendValue(b, v)
	}
}

func (f *LogrusTextFormatter) needsQuoting(text string) bool {
	if f.QuoteEmptyFields && len(text) == 0 {
		return true
	}
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.' || ch == '_' || ch == '/' || ch == '@' || ch == '^' || ch == '+') {
			return true
		}
	}
	return false
}

func (f *LogrusTextFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString(key)
	b.WriteByte('=')
	f.appendValue(b, value)
}

func (f *LogrusTextFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	if !f.needsQuoting(stringVal) {
		b.WriteString(stringVal)
	} else {
		b.WriteString(fmt.Sprintf("%q", stringVal))
	}
}
