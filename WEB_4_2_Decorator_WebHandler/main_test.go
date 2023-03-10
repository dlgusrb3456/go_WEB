package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPage(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestDecoHandler(t *testing.T) { //로그가 잘 찍혔는지 확인
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	buf := &bytes.Buffer{}
	log.SetOutput(buf) //sets the output destination for the standard logger

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	r := bufio.NewReader(buf)
	line, _, err := r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Started")

	line, _, err = r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Completed")
}
