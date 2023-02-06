package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// membuat set cookie
func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "x-pzn-name"
	cookie.Value =request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer,cookie)

	fmt.Fprint(writer,"Sucess Create Cookie",cookie.Value)
}

//membuat get cookie
func GetCookie(writer http.ResponseWriter,request *http.Request)  {
	cookie, err := request.Cookie("x-pzn-name")
	if err != nil {
		fmt.Fprint(writer,"No Cookie")
	}else{
		name := cookie.Value
		fmt.Fprintf(writer, "Hello  %s", name)
	}
}

// mencoba menjalankan set dan get cookie yang telah dibuat kedalam server
func TestCookie(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie",SetCookie)
	mux.HandleFunc("/get-cookie",GetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

// unit test set cookie
func TestSetCookie(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080/?name=dede",nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder,request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies{
		fmt.Printf("Cookie %s:%s \n",cookie.Name,cookie.Value)
	}
}

func TestGetCookie(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080/",nil)

	cookie := new(http.Cookie)
	cookie.Name = "x-pzn-name"
	cookie.Value = "eko"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder,request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Print(bodyString)
}