package main
import "net/http"

func main {
	//http.ListenAndServe builds a webserver and when it receives a request it will hand it over to http.handler.
	//http.FileServer implements the handler interface
	http.ListenAndServe(":8080",http.FileServer("."))
}
