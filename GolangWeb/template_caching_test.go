package GolangWeb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates1 embed.FS

var myTemplates = template.Must(template.ParseFS(templates1, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template Caching")
}

func TestTemplateFunctionCaching(t *testing.T) {
	reqwuest := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, reqwuest)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
