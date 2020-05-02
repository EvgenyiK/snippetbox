package main

import(
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello from Snippetbox"))
}

func main()  {
	
}