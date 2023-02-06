package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

//middleware
type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	fmt.Println("Before execute Hanlder")
	middleware.Handler.ServeHTTP(writer,request)
	fmt.Println("After Execute Handler")
}

//error handler middleware
type ErrorHandler struct{
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer,"Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer,request)
}

func TestMiddleware(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/",func (writer http.ResponseWriter, request *http.Request)  {
		fmt.Println("Handler Execute")
		fmt.Fprintf(writer, "Hello Middleware")	
	})
	mux.HandleFunc("/foo",func (writer http.ResponseWriter, request *http.Request)  {
		fmt.Println("foo Execute")
		fmt.Fprintf(writer, "Hello Foo")	
	})
	mux.HandleFunc("/panic",func (writer http.ResponseWriter, request *http.Request)  {
		fmt.Println("foo Execute")
		panic("ups")
	})

	//3. logmiddleware ekseskusi mux
	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	//2. errorHandler ekseskusi logMiddleware
	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	//1. server eksekusi errorHandler
	server := http.Server{
		Addr: "localhost:8080",
		Handler: errorHandler,
	}
	err:= server.ListenAndServe()
	
	if err != nil{
		panic(err)
	}
}

