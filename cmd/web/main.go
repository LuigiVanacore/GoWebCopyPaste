package main

import (
	"flag"
	"github.com/LuigiVanacore/GoWebCopyPaste/controllers"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/snippet", controllers.ShowSnippet)
	mux.HandleFunc("/snippet/create", controllers.CreateSnippet)

	fileServer := http.FileServer(http.Dir("././ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
