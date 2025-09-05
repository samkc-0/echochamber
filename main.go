package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("starting app..")
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET / was requested!")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("exiting app..")
}
