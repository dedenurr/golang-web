package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/layout.gohtml",
		"./templates/footer.gohtml",
		
	))
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title":"Template Layout",
		"Name":"Dede",
	})
}

func TestTemplateLayout(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder,request)

	response :=recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}