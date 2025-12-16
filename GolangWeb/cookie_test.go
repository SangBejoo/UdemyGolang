package GolangWeb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Ayub-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/" // ini berarti cookie untuk semuanya

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success create cookies")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-Ayub-Name")
	if err != nil {
		fmt.Fprint(writer, "Error get cookie from request")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", setCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Ayub", nil)
	recorder := httptest.NewRecorder()
	setCookie(recorder, request)

	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Printf("Cookie name: %s:%s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-Ayub-Name"
	cookie.Value = "ayub"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()
	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Body)
	fmt.Println(string(body))
}
