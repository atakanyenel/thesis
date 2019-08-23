package main

import "net/http"

func main() {
	_, err := http.Get("http://localhost:8080/count?payload=atakan")
	if err != nil {
		panic(err)
	}
}
