package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("vim-go")

	port := os.Getenv("PORT")
	if port == "" {
		panic("must specify PORT env var")
	}

	http.HandleFunc("/healthz", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	go func() {
		http.ListenAndServe(":"+port, nil)
	}()

	log.Fatalln(http.ListenAndServe(":8888", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	})))
}
