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
	file, _ := os.Open(path) //opens the named file for reading
	defer file.Close()

	os.RemoveAll("./uploads") // 해당 경로의 모든것 삭제 ,removes path and any children it contains

	buf := &bytes.Buffer{}             // bytes 패키지의 Buffer 타입 구조체
	writer := multipart.NewWriter(buf) // NewWriter returns a new multipart Writer with a random boundary, writing to w. w에 적을 writer를 반환한다.
	// MIME형식으로 파일을 웹상에서 주고 받음. 바이너리 -> 텍스트(인코딩) , 텍스트 -> 바이너리(디코딩) 의 과정을 거침. 이 MIME 형식을 맞추기 위해 multipart를 사용함
	w, err := writer.CreateFormFile("upload_file", filepath.Base(path)) //It creates a new form-data header with the provided field name and file name.
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

	uploadFile, _ := os.Open(uploadFilePath) //특정 경로 파일 열기,opens the named file for reading. If successful, methods on the returned file can be used for reading; the associated file descriptor has mode O_RDONLY.
	originFile, _ := os.Open(path)

	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}

	uploadFile.Read(uploadData) //연거 읽기,reads up to len(b) bytes from the File and stores them in b. It returns the number of bytes read and any error encountered. At end of file, Read returns 0, io.EOF
	originFile.Read(originData)

	assert.Equal(uploadData, originData)
}
