package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 1. Request Header

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost/",nil)
	request.Header.Add("Content-Type","Aplication/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)

}

// 2. Response Header
func ResponseHeader(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Add("X-Powered-By","Dede Nurrahman")
	fmt.Fprint(writer,"OK oh yeas")
}

func TestResponseHeader(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder,request)

	// get header
	poweredBy := recorder.Header().Get("x-powered-by")
	fmt.Println(poweredBy)

	// get body
	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)

}