package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile, header, err := r.FormFile("upload_file") // input formfile
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()
	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename) //업로드 된 파일의 파일 경로
	file, err := os.Create(filepath)
	defer file.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}

	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

func main() {
	http.HandleFunc("/uploads", uploadsHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":3000", nil)
}
