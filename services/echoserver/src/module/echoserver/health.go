package main

import (
	"fmt"
	"net/http"
)

func healthReady(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok\n")
}

func healthAlive(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok\n")
}
