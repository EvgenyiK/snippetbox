package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// flag will be stored in the addr variable at runtime.
	addr:=flag.String("addr",":4000","HTTP network address")
	//encountered during parsing the application will be terminated.
	flag.Parse()


	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	// Create a file server which serves files out of the "./ui/static" directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
