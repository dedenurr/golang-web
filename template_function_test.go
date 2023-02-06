package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

// 1. template function
type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name Is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`)) //SayHello => data Eko(myPage.Name), Budi => mengisi data name pada SayHello

	t.ExecuteTemplate(writer,"FUNCTION", MyPage{
		Name:"Eko",
	})
}

func TestTemplateFunction(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder,request)

	response :=recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

// 2. Template Global Function
func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`)) // len tidak menggunakan "." didepannya karena dia sifatnya global

	t.ExecuteTemplate(writer,"FUNCTION", MyPage{
		Name:"Tutorial Go-Lang",
	})
}

func TestTemplateFunctionGlobal(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder,request)

	response :=recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

// 3. Menambah Template Global Function
func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request)  {
	// buat template baru dengan nama FUNCTION
	t := template.New("FUNCTION")
	// registrasi/Isikan func berupa map interface untuk merubah huruf jadi uppercase
	t = t.Funcs(map[string]interface{}{
		"upper": func (value string) string  {
			return strings.ToUpper(value)	
		},
	})

	// parsing data
	t = template.Must(t.Parse(`{{upper .Name}}`))

	// ekseskusi value
	t.ExecuteTemplate(writer, "FUNCTION",MyPage{
		Name:"zahra",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder,request)

	response :=recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

// 4. Template Pipelines Function
func TemplateFunctionPipelines(writer http.ResponseWriter, request *http.Request)  {
	// buat template baru dengan nama FUNCTION
	t := template.New("FUNCTION")
	// registrasi/Isikan func berupa map interface untuk merubah huruf jadi uppercase
	t = t.Funcs(map[string]interface{}{
		"sayHello" : func (value string)string  {
			return "hello " + value	
		},
		"upper": func (value string) string  {
			return strings.ToUpper(value)	
		},
	})

	// parsing data mulai dari funciont global sayHello lanjut ke upper
	t = template.Must(t.Parse(`{{ sayHello .Name | upper}}`))

	// ekseskusi value
	t.ExecuteTemplate(writer, "FUNCTION",MyPage{
		Name:"zahra",
	})
}

func TestTemplateFunctionPipelines(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipelines(recorder,request)

	response :=recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}