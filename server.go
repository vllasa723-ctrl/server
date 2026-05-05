package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	
	go func() {
		err := http.ListenAndServe("127.0.0.1:8080", nil)
		if err != nil {
			log.Printf("HTTP error: %v", err)
		}
	}()
	
	// Оберни в log.Fatal, чтобы видеть причину смерти!
	log.Fatal(http.ListenAndServeTLS("127.0.0.1:8443", "cert.pem", "key.pem", nil))
}
