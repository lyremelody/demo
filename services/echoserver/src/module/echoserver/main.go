package main

import (
	"fmt"
	"net/http"

	"github.com/lyremelody/demo/utils/config"
)

func main() {
	serviceinfo, err := config.GetServiceInfo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	address := fmt.Sprintf(":%d", serviceinfo.Port)

	http.HandleFunc("/api/echoserver/v1/nowtime", echoNowTime)
	http.HandleFunc("/api/echoserver/v1/sleep/nowtime", sleepAndEchoNowTime)

	http.HandleFunc("/health/ready", healthReady)
	http.HandleFunc("/health/alive", healthAlive)

	http.ListenAndServe(address, nil)
}
