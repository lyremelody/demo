package main

import (
	"fmt"
	"net/http"

	"github.com/lyremelody/demo/module/echo"
)

func echoNowTime(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Now: %v\n", echo.GetNowTime())
}
