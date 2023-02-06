package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// 1. query parameter
func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == ""{
		fmt.Fprint(writer,"Hello")
	}else{
		fmt.Fprintf(writer,"Hello %s",name)
	}
	

}
func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET","http://localhost:8080/hello?name=Dede",nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder,request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

// 2. multiple query parameter
func MultipleParameter(writer http.ResponseWriter, request *http.Request)  {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer,"%s %s", firstName, lastName)
	
}

func TestQueryMultipleParameter(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost/?first_name=Dede&last_name=Nurrahman",nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder,request)

	response := recorder.Result()
	body, _:= io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

//3. multiple value query parameter
func MultipleValueParameter(writer http.ResponseWriter,request *http.Request)  {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprintln(writer,strings.Join(names," "))
}

func TestMultipleValueParameter(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost/?name=dede&name=Zahra&name=adel",nil)
	recorder := httptest.NewRecorder()

	MultipleValueParameter(recorder,request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}