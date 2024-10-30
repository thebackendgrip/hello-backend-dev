package main

import "net/http"

func main() {
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
	})
}
