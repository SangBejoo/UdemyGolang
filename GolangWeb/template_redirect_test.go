package GolangWeb

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func RedirectFrom(writer http.ResponseWriter, req *http.Request) {
	http.Redirect(writer, req, "https://www.youtube.com/", http.StatusTemporaryRedirect)
}

func RedirectOut(writer http.ResponseWriter, req *http.Request) {
	http.Redirect(writer, req, "https://lmarena.ai/c/019b52e6-2d04-7751-b1a8-4c7c60d852f2", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
