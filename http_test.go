package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func HeloHandler(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprint(writer, "Hello World")

}

// membuat test tanpa harus menjalankan server
func TestHelloHandler(t *testing.T) {
	// buat request dan recorder
	request := httptest.NewRequest("GET","http://localhost:8080/hello",nil)
	recorder := httptest.NewRecorder()

	// panggil handler
	HeloHandler(recorder,request)

	// cek body responnya
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}