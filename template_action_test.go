package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// 1. Template action dengan if else
func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Template Action Data",
		Name:  "",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

// 2. Template action dengan if else operator
func TemplateActionOperator(writer http.ResponseWriter,request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	t.ExecuteTemplate(writer,"comparator.gohtml",map[string]interface{}{
		"Title": "Template Action Data",
		"FinalValue" : 50,
	})
}


func TestTemplateActionOperator(t *testing.T)  {
	request := httptest.NewRequest("GET","http:/localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

// 3. Template action dengan if else & range
func TemplateRange(writer http.ResponseWriter,request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(writer,"range.gohtml",map[string]interface{}{
		"Hobbies" : []string{
			"Gaming", "Futsal", "Coidng",
		},
	})
}

func TestTemplateRange(t *testing.T)  {
	request := httptest.NewRequest("GET","http:/localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateRange(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

// 4. Template action dengan with
func TemplateWith(writer http.ResponseWriter,request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))

	t.ExecuteTemplate(writer,"address.gohtml",map[string]interface{}{
		"Title":"Template Action With",
		"Name":"dede",
		"Address": map[string]interface{}{
			"Street":"Jalan Saturnus",
			"City":"Pluto",
		},
	})
}

func TestTemplateWith(t *testing.T)  {
	request := httptest.NewRequest("GET","http:/localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}