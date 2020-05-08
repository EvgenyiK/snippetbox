package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct{
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	// flag will be stored in the addr variable at runtime.
	addr:=flag.String("addr",":4000","HTTP network address")
	//encountered during parsing the application will be terminated.
	flag.Parse()

	// Create a logger for writing error messages in the same way, but use stderr as
	infoLog:=log.New(os.Stdout,"INFO\t",log.Ldate|log.Ltime|log.Lshortfile)
	errorLog:=log.New(os.Stderr,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of application containing the dependencies.
	app:=&application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)
	// Create a file server which serves files out of the "./ui/static" directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Initialize a new http.Server struct. We set the Addr and Handler
	srv:=&http.Server{
		Addr:	*addr,
		ErrorLog: errorLog,
		Handler: mux,
	}
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
