package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Membuat web dinamis dengan Template
// 1. Template String
func SimpleHTML(writer http.ResponseWriter,request *http.Request) {
	templateText :=  "<html><body>{{.}}</body></html>"
	// t, err:= template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }
	t := template.Must(template.New("SIMPLE").Parse(templateText))
	t.ExecuteTemplate(writer,"SIMPLE","HELLO HTML Template Dede")
}

func TestTemplateString(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder,request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)

}

//2. Template dari File
func SimpleHTMLFile(writer http.ResponseWriter,request *http.Request)  {
	t:= template.Must(template.ParseFiles("./templates/simple.gohtml"))

	t.ExecuteTemplate(writer,"simple.gohtml","Hello HTML TEMPLATE")
}

func TestSimpleHTMLTemplate(t *testing.T)  {
	request := httptest.NewRequest("GET","http:/localhost",nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

//3. Template directory
func TemplateDirectory(writer http.ResponseWriter,request *http.Request)  {
	t:= template.Must(template.ParseGlob("./templates/*.gohtml"))

	t.ExecuteTemplate(writer,"simpleow.gohtml","Hello HTML TEMPLATE")
}

func TestTemplateDirectory(t *testing.T)  {
	request := httptest.NewRequest("GET","http:/localhost",nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

//4. Template Embed


func TemplateEmbed(writer http.ResponseWriter,request *http.Request)  {
	t:= template.Must(template.ParseFS(templates,"templates/*.gohtml"))

	t.ExecuteTemplate(writer,"simple.gohtml","Hello HTML TEMPLATE")
}

func TestTemplateEmbed(t *testing.T)  {
	request := httptest.NewRequest("GET","http:/localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}