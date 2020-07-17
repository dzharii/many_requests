package main

import (
	"log"
	"net/http"
	"time"
)

func makeRequest() {
	resp, err := http.Get("http://localhost:3000/sleep")
	if err != nil {
		log.Print(err)
		return
	}

	log.Print(resp.Status)

	defer resp.Body.Close()
}

func main() {
	for i := 0; i < 100000; i++ {
		log.Printf("making request %d", i)
		time.Sleep(time.Microsecond)
		go makeRequest()
	}
	time.Sleep(time.Hour)
}
