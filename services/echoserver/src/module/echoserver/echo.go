package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lyremelody/demo/module/echo"
)

func sleepAndEchoNowTime(w http.ResponseWriter, req *http.Request) {
	time.Sleep(10 * time.Second)
	fmt.Fprintf(w, "Now: %v\n", echo.GetNowTime())
}

func echoNowTime(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Now: %v\n", echo.GetNowTime())
}
