package log

import (
	"log/slog"
	"net/http"
	"os"
	"thewindear/test-go-web/src/core"
	"time"
)

var Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func BuildRequestLogger(r *http.Request) *slog.Logger {
	now := time.Now()
	reqId := now.UnixNano()
	return Logger.With(slog.Group(
		"request",
		"url",
		r.URL.Path,
		"query",
		r.URL.RawQuery,
		"method",
		r.Method,
		"request_id",
		reqId,
		"version",
		core.ApiVersion,
	))
}
