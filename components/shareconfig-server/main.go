package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "get only", http.StatusMethodNotAllowed)
			return
		}
		username, password, authOK := r.BasicAuth()

		if authOK == false {
			http.Error(w, "Not authorized", 401)
			return
		}
		if username != "username" && password != "password" {
			http.Error(w, "Not authorized", 401)
			return
		}
		http.ServeFile(w, r, "shareconfig.txt")
	})

	log.Print("Starting server on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
