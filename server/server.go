package main

import (
	"log"
	"net/http"
	"time"
)

func sleep(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Got request")
	time.Sleep(time.Minute)
	writer.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/sleep", sleep)
	http.ListenAndServe(":3000", nil)
}
