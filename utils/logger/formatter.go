package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"sort"
	"time"
)

const defaultTimestampFormat = "2006-01-02T15:04:05.9999Z07:00"

type Clock interface {
	Now() time.Time
	Since(t time.Time) time.Duration
}

type realClock struct{}

func (realClock) Now() time.Time                  { return time.Now() }
func (realClock) Since(t time.Time) time.Duration { return time.Since(t) }

// textFormatter formats logs into text
type textFormatter struct{}

var clock Clock = new(realClock)

// Setting Clock is only used in test purposes
func setClock(newClock Clock) {
	clock = newClock
}

// Format renders a single log entry
func (f *textFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Time = clock.Now()

	keys := make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	f.appendKeyValue(b, "level", entry.Level.String())
	f.appendKeyValue(b, "time", entry.Time.Format(defaultTimestampFormat))
	if entry.Message != "" {
		f.appendKeyValue(b, "msg", entry.Message)
	}
	for _, key := range keys {
		f.appendKeyValue(b, key, entry.Data[key])
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *textFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString(key)
	b.WriteByte('=')
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}
	b.WriteString(stringVal)
}
