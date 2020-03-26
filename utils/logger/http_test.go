package logger

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggerMiddlewareStatusOk(t *testing.T) {
	setClock(new(mockClock))
	var buf bytes.Buffer
	httpLogger.SetOutput(&buf)
	expected := "level=info time=" + expectedTime + " duration=100ms method=GET path=/test size=0B status=200\n"
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(Middleware)

	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
		return

	})
	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)

	r.ServeHTTP(resp, c.Request)

	output := buf.String()
	if output != expected {
		t.Error(output)
	}
}

func TestLoggerMiddlewareStatusFound(t *testing.T) {
	setClock(new(mockClock))
	var buf bytes.Buffer
	httpLogger.SetOutput(&buf)
	expected := "level=info time=" + expectedTime + " duration=100ms method=GET path=/test size=0B status=302\n"
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(Middleware)

	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusFound)
		return

	})
	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)

	r.ServeHTTP(resp, c.Request)

	output := buf.String()
	if output != expected {
		t.Error(output)
	}
}

func TestLoggerMiddlewareStatusNotFound(t *testing.T) {
	setClock(new(mockClock))
	var buf bytes.Buffer
	httpLogger.SetOutput(&buf)
	expected := "level=warning time=" + expectedTime + " duration=100ms method=GET path=/test size=0B status=404\n"
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(Middleware)

	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusNotFound)
		return

	})
	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)

	r.ServeHTTP(resp, c.Request)

	output := buf.String()
	if output != expected {
		t.Error(output)
	}
}

func TestLoggerMiddlewareStatusInternalServerError(t *testing.T) {
	setClock(new(mockClock))
	var buf bytes.Buffer
	httpLogger.SetOutput(&buf)
	expected := "level=error time=" + expectedTime + " duration=100ms method=GET path=/test size=0B status=500\n"
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(Middleware)

	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusInternalServerError)
		return

	})
	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)

	r.ServeHTTP(resp, c.Request)

	output := buf.String()
	if output != expected {
		t.Error(output)
	}
}

// level=error time="2020-10-15T14:15:45Z" duration=100ms method=GET path=/test req-id=27af7c8b size=0B status=500
// level=error time="2020-10-15T14:15:45+03:00" duration=100ms method=GET path=/test req-id=27af7c8b size=0B status=500
