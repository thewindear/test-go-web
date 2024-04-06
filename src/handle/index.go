package handle

import (
	"fmt"
	"net/http"
	"thewindear/test-go-web/src/core"
	"thewindear/test-go-web/src/log"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	reqId := now.UnixNano()
	nowTime := now.Format(time.RFC3339)
	log.BuildRequestLogger(r).Info("")
	result := fmt.Sprintf(
		"verion: %s, path: %s, request time: %s, request_id: %d\n%s", core.ApiVersion, r.URL.Path, nowTime, reqId, r.URL.RawQuery)
	w.Header().Set("X-Request-ID", fmt.Sprintf("%d", reqId))
	w.Header().Set("Version", core.ApiVersion)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(result))
}
