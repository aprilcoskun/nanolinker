package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"time"

	"github.com/sirupsen/logrus"
)

const log1000 = 6.907755278982137

var byteSizes = [4]string{"B", "kB", "MB", "GB"}
var httpLogger *logrus.Logger

// Http Logger Middleware
func Middleware(c *gin.Context) {
	start := time.Now()
	c.Next()

	statusCode := c.Writer.Status()

	fields := logrus.Fields{
		"status":   c.Writer.Status(),
		"method":   c.Request.Method,
		"path":     c.Request.URL.String(),
		"duration": clock.Since(start).String(),
	}

	dataLength := c.Writer.Size()
	if dataLength <= 0 {
		fields["size"] = "0B"
	} else {
		fields["size"] = humanizeBytes(uint64(dataLength))
	}

	entry := httpLogger.WithFields(fields)
	if len(c.Errors) > 0 {
		entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
	} else {
		if statusCode > 499 {
			entry.Error()
		} else if statusCode > 399 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}

func humanizeBytes(s uint64) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}

	e := math.Floor(math.Log(float64(s)) / log1000)
	val := math.Floor(float64(s)/math.Pow(1000, e)*10+0.5) / 10
	f := "%.0f%s"
	if val < 10 {
		f = "%.1f%s"
	}

	return fmt.Sprintf(f, val, byteSizes[int(e)])
}
