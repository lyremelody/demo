package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/api/echoserver/v1/echo/nowtime", echoNowTime)
	http.HandleFunc("/health/ready", healthReady)
	http.HandleFunc("/health/alive", healthAlive)
	http.ListenAndServe(":80", nil)
}
