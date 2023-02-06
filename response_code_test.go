package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == ""{
		writer.WriteHeader(/* 400 */http.StatusBadRequest)//bad request
		fmt.Fprint(writer,"Name Is Empty")
	}else{
		writer.WriteHeader(200)
		fmt.Fprintf(writer,"Hi %s",name)
	}

}
func TestResponseCode(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost:8080/",nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodySring := string(body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(bodySring)

}