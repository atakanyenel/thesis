package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func counter(w http.ResponseWriter, r *http.Request) {
	incomingTime := time.Now().Unix()
	incominMessage := r.URL.Query()["payload"][0]
	fmt.Printf("%v,%s\n", incomingTime, incominMessage)
	fmt.Fprint(w, "OK")
}
func main() {
	http.HandleFunc("/count", counter)

	log.Print("Starting server on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func startLogging(incomingTime string, message string, count uint64) {
	go func(incomingTime string, message string, count uint64) {
		fmt.Printf("%s,%s,%s\n", incomingTime, message, count)
		time.Sleep(1 * time.Second)
	}(incomingTime, message, count)
}
