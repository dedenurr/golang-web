package golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Display Form Upload (pertama menampilkan form upload dan menggunakan myTemplates dari template chacing untuk membaca file gohtml dari folder templates)
func UploadForm(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w,"upload.form.gohtml",nil)
}

// Upload Handler (semua file di handle/ditangani oleh kode ini)
func Upload(w http.ResponseWriter, r *http.Request)  {
	// menangkap data file
	file, fileHeader, err := r.FormFile("file")	
	if err != nil {
		panic(err)
	}

	// buat lokasi penyimpanana file
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	//simpan hasil upload (file) ke dalam (fileDestination)
	_,err = io.Copy(fileDestination,file)
	if err != nil {
		panic(err)
	}

	//jika data bukan berupa file 
	name := r.PostFormValue("name")

	//eksekusi file/render
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml",map[string]interface{}{
		"Name":name,
		"File": "/static/" + fileHeader.Filename,
	})


}


// Server Upload (coba di tes dengan kode ini menggunakan mux server)
func TestUploadForm(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/",UploadForm)
	mux.HandleFunc("/upload",Upload)
	mux.Handle("/static/",http.StripPrefix("/static",http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err :=	server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}



//go:embed resources/code.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T)  {
	// tempat penyimpan
	body := new(bytes.Buffer)
	// data yang dikirimkan ke body, format dari form file (multpart)
	writer := multipart.NewWriter(body)
	// isi field
	writer.WriteField("name","Dede Nurrahman")

	// mengirimkan file  yg akan di upload
	file, _ := writer.CreateFormFile("file", "CONTOHUPLOAD.png")
	// file test dari go embed
	file.Write(uploadFileTest)
	// agar tidak ada memori yang menggantung
	writer.Close()
	
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}