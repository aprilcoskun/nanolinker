package utils

import (
	"github.com/nanmu42/gzip"
)

var GzipMiddleware = gzip.NewHandler(gzip.Config{
	// gzip compression level to use
	CompressionLevel: 6,
	// minimum content length to trigger gzip, 10kb
	MinContentLength: 1024 * 10,
	// Only compress static files
	RequestFilter: []gzip.RequestFilter{
		gzip.NewCommonRequestFilter(),
		gzip.NewExtensionFilter([]string{"", ".css", ".js"}),
	},
}).Gin
