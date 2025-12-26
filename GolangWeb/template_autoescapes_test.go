package GolangWeb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscapes(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escapes",
		"Body":  "<h1>Hello Ayub Subagiya<script> anda di hacel</script>w</h1>",
	})
}

func TemplateAutoEscapesDissabled(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escapes",
		"Body":  template.HTML("<h1>Hello Ayub Subagiya<script> anda di hacel</script>w</h1>"),
	})
}

func TemplateXss(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escapes",
		"Body":  template.HTML(request.URL.Query().Get("body")),
	})
}

func TestTemplateXss(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>Alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXss(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXssServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXss),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
