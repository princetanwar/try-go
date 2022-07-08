package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
		fmt.Println("request received ")
	}).Methods("get")

	srv := &http.Server{Handler: router, Addr: ":8001"}

	log.Fatal(srv.ListenAndServe())
}
