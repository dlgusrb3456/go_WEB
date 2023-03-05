package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)

	path := "C:/Users/dlgus/OneDrive/바탕 화면/라파.jpg"
	file, _ := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads") // 해당 경로의 모든것 삭제

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	w, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)

	io.Copy(w, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType()) //해당 request가 어떤 type인지 알려줘야 함

	uploadsHandler(res, req)

	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath) //uploadFilePath의 상태를 return하고 error가 있는지 확인. 없으면 파일이 존재하는거임.

	assert.NoError(err)

	uploadFile, _ := os.Open(uploadFilePath) //특정 경로 파일 열기
	originFile, _ := os.Open(path)

	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}

	uploadFile.Read(uploadData) //연거 읽기
	originFile.Read(originData)

	assert.Equal(uploadData, originData)
}
