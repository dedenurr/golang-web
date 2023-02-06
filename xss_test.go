package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 1
// auto escape adalah fitur di golang untuk menangani XSS Cross Site Scripting
func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml",map[string]interface{}{
		"Title":"Template Auto Escape",
		"Body": "<p>Ini adalah body<p>alert('Anda Di Heck')</script></p>",
	})
}

func TestTemplateAutoEscape(t *testing.T)  {
	request := httptest.NewRequest("GET","http:/localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestTemplateAutoEscapeServer(t *testing.T){
	server := http.Server{
		Addr:"localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// 2
// template html bisa di matikan dengan cara
/* template.HTML , jika ini adalah data html
template.CSS, jika ini adalah data css
template.JS, jika ini adalah data javascript
 */
func TemplateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml",map[string]interface{}{
		"Title":"Template Auto Escape",
		"Body": template.HTML("<h1>Ini adalah body</h1>"),
	})
}


func TestTemplateAutoEscapeDisabled(t *testing.T)  {
	request := httptest.NewRequest("GET","http:/localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T){
	server := http.Server{
		Addr:"localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


// Kasus XSS
func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml",map[string]interface{}{
		"Title":"Template Auto Escape",
		"Body": template.HTML(request.URL.Query().Get("body")),
	})
}


func TestTemplateXSS(t *testing.T)  {
	request := httptest.NewRequest("GET","http:/localhost:8080/?body=<p>alert</p>",nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestTemplateXSSServer(t *testing.T){
	server := http.Server{
		Addr:"localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
