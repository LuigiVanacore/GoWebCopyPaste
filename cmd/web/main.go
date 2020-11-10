package main

import (
	"github.com/LuigiVanacore/GoWebCopyPaste/controllers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/snippet", controllers.ShowSnippet)
	mux.HandleFunc("/snippet/create", controllers.CreateSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
