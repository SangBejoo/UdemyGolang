package GolangWeb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionPerbandingan(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/perbandingan.gohtml"))
	t.ExecuteTemplate(writer, "perbandingan.gohtml", map[string]interface{}{
		"Title":      "Template Action Perbandingan",
		"FinalValue": 80,
	})

}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionPerbandingan(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
