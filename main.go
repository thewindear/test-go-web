package main

import (
	"fmt"
	"net/http"
	"thewindear/test-go-web/src/handle"
	"thewindear/test-go-web/src/log"
)

func main() {
	http.HandleFunc("/", handle.Index)
	fmt.Println("Listening on port 8080")
	log.Logger.Error("server run error: %s", http.ListenAndServe(":8080", nil))
}
