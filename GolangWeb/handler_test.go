package GolangWeb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writer, "Ini Ayub Subagiya Test Handler")
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// serve Mux
func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Ini Ayub Subagiya Test Handler deengan mux")
	})
	mux.HandleFunc("/Hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi ini Ayub Subagiya Test Handler deengan mux")
	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hmmm apakah bisa")
	})
	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Image thumbnails test")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/Hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Ini Ayub Subagiya Test Handler deengan mux")
	})
	mux.HandleFunc("/Add/", func(writer http.ResponseWriter, request *http.Request) {
		var data struct {
			X int `json:"x"`
			Y int `json:"y"`
		}
		if request.Method == http.MethodPost {
			if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
				http.Error(writer, "Invalid JSON", http.StatusBadRequest)
				return
			}
		} else if request.Method == http.MethodGet {
			xStr := request.URL.Query().Get("x")
			yStr := request.URL.Query().Get("y")
			var err error
			data.X, err = strconv.Atoi(xStr)
			if err != nil {
				http.Error(writer, "Invalid x parameter", http.StatusBadRequest)
				return
			}
			data.Y, err = strconv.Atoi(yStr)
			if err != nil {
				http.Error(writer, "Invalid y parameter", http.StatusBadRequest)
				return
			}
		} else {
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		result := Add(data.X, data.Y)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(map[string]int{"result": result})
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func Add(x, y int) int {
	return x + y
}
