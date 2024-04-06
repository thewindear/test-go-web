package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

const ApiVersion = "0.2"

func BuildLogger(r *http.Request) *slog.Logger {
	now := time.Now()
	reqId := now.UnixNano()
	return logger.With(slog.Group(
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
		ApiVersion,
	))
}

func Index(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	reqId := now.UnixNano()
	nowTime := now.Format(time.RFC3339)
	BuildLogger(r).Info("")
	result := fmt.Sprintf(
		"path: %s, request time: %s, request_id: %d\n%s", r.URL.Path, nowTime, reqId, r.URL.RawQuery)
	w.Header().Set("X-Request-ID", fmt.Sprintf("%d", reqId))
	w.Header().Set("Version", ApiVersion)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(result))
}

func main() {
	http.HandleFunc("/", Index)
	fmt.Println("Listening on port 8080")
	logger.Error("server run error: %s", http.ListenAndServe(":8080", nil))
}
