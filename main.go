package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/validate", validate)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("http server shut down: %v\n", err)
	}
}

func validate(w http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("unable to read request bytes: %v\n", err)
	}

	log.Println("Request: ", string(b))

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"msg": "ok"}`))
}
