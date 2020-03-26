package logger

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"
)

type mockClock struct{}

func (mockClock) Now() time.Time { return time.Date(2020, 10, 15, 14, 15, 45, 0, time.Local) }
func (mockClock) Since(_ time.Time) time.Duration {
	duration, _ := time.ParseDuration("100ms")
	return duration
}

var expectedTime = time.Date(2020, 10, 15, 14, 15, 45, 0, time.Local).Format(defaultTimestampFormat)

func captureOutputAndMockTime(f func()) string {
	setClock(new(mockClock))
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func TestInfo(t *testing.T) {
	output := captureOutputAndMockTime(func() { Info("info log") })
	expected := "level=info time=" + expectedTime + " msg=info log\n"
	if expected != output {
		t.Error("info log fail")
	}
}

func TestWarn(t *testing.T) {
	output := captureOutputAndMockTime(func() { Warn("warn log") })
	expected := "level=warning time=" + expectedTime + " msg=warn log\n"
	if expected != output {
		t.Error("warning log fail")
	}
}

func TestError(t *testing.T) {
	output := captureOutputAndMockTime(func() { Error("error log") })
	expected := "level=error time=" + expectedTime + " msg=error log\n"
	if expected != output {
		t.Error("error log fail")
	}

}
