package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	go http.ListenAndServe("127.0.0.1:8080", nil)
	http.ListenAndServeTLS("127.0.0.1:8443", "cert.pem", "key.pem", nil)
}
