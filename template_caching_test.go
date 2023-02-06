package golang_web

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
var templates embed.FS

// template di simpan diluar agar lebih cepat untuk prosesnya jangan disimpan didalam function TemplateCaching

var myTemplates = template.Must(template.ParseFS(templates,"templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request)  {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateCaching(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder,request)

	response :=recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
