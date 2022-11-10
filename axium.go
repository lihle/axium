package main

import (
	"axium/handler"
	"axium/storage"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

//The main function
func main() {
	fmt.Println("Hello-World")
	//
	err := storage.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := mux.NewRouter()

	//Pages
	r.HandleFunc("/", handler.Viewlogin)

	//Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	port := ":3001"
	fmt.Println("http://localhost" + port)
	fmt.Println()
	http.ListenAndServe(port, r)
}
