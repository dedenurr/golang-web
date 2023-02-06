package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// 1. Template Data dengan Map
func TemplateDataMap(writer http.ResponseWriter,request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer,"name.gohtml",map[string]interface{}{
		"Title" : "Template Data Map",
		"Name" : "Dede",
		"Address" : map[string]interface{}{
			"Street" : "Jalan Ali Muchtar 2",
		},
	})
}

func TestTemplateDataMap(t *testing.T)  {
	request := httptest.NewRequest("GET","http:/localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

//  2. Template Data dengan Struct
type Address struct{
	Street string
}


type Page struct{
	Title string
	Name string
	Address Address
}


func TemplateDataStruct(writer http.ResponseWriter,request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name: "Dede Zahra",
		Address: Address{
			Street: "Jalan Ali Muchtar",
		},
	})
}

func TestTemplateDataStruct(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
