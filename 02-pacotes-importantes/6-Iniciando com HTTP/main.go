package main

import "net/http"

func main() {

	http.HandleFunc("/", BuscaCEP)
	/*	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello, world!"))
	})*/

	http.ListenAndServe(":8080", nil)
}

func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, ,World!"))
}
