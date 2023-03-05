package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	indexHandler(res, req) // "/" 경로 핸들러
	assert.Equal(http.StatusOK, res.Code, "Failed!!")

	// if res.Code != http.StatusOK { // res의 status를 확인
	// 	t.Fatal("Failed!!", res.Code)
	// }

	// res body 읽어와서 확인하기
	data, _ := ioutil.ReadAll(res.Body) //Body는 Buffer값이라 바로 읽어올 수 없음. 그래서 ioutil을 이용해 버퍼의 내용을 전부 읽어와 data에 저장할 것임
	assert.Equal("hello world", string(data), "data failed")

}

func TestBarPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar2", nil)
	//req := httptest.NewRequest("GET", "/", nil) // 이렇게 해도 통과가 됨. target이 "/" 지만 barhandler를 직접 호출했기 때문에 target이 제대로 적용되고 있지 않음. => 이 문제를 해결하기 위해서 mux를 사용해줘야함

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// barHandler(res, req) // "/" 경로 핸들러

	assert.Equal(http.StatusOK, res.Code, "Failed!!")

	// if res.Code != http.StatusOK { // res의 status를 확인
	// 	t.Fatal("Failed!!", res.Code)
	// }

	// res body 읽어와서 확인하기
	data, _ := ioutil.ReadAll(res.Body) //Body는 Buffer값이라 바로 읽어올 수 없음. 그래서 ioutil을 이용해 버퍼의 내용을 전부 읽어와 data에 저장할 것임
	assert.Equal("hello world!", string(data), "data failed")

}

func TestBarPathHandler_withname(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar2?name=asdf", nil)
	//req := httptest.NewRequest("GET", "/", nil) // 이렇게 해도 통과가 됨. target이 "/" 지만 barhandler를 직접 호출했기 때문에 target이 제대로 적용되고 있지 않음. => 이 문제를 해결하기 위해서 mux를 사용해줘야함

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// barHandler(res, req) // "/" 경로 핸들러

	assert.Equal(http.StatusOK, res.Code, "Failed!!")

	// if res.Code != http.StatusOK { // res의 status를 확인
	// 	t.Fatal("Failed!!", res.Code)
	// }

	// res body 읽어와서 확인하기
	data, _ := ioutil.ReadAll(res.Body) //Body는 Buffer값이라 바로 읽어올 수 없음. 그래서 ioutil을 이용해 버퍼의 내용을 전부 읽어와 data에 저장할 것임
	assert.Equal("hello asdf!", string(data), "data failed")

}
