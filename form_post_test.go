package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("firstName")
	lastName := request.PostForm.Get("lastName")

	fmt.Fprintf(writer,"%s %s ",firstName, lastName)
	fmt.Fprint(writer,"|| OK Form Post Berhasil terkirim")
}

func TestFormPost(t *testing.T)  {
	requestBody := strings.NewReader("firstName=Dede&lastName=Nurrahman")
	request := httptest.NewRequest(/*http.MethodPost*/"POST","http://localhost:8080",requestBody)
	request.Header.Add("Content-Type","application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}