package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	res, err := http.Get(ts.URL)

	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body) //Body는 Buffer값이라 바로 읽어올 수 없음. 그래서 ioutil을 이용해 버퍼의 내용을 전부 읽어와 data에 저장할 것임
	assert.Equal("Hello World", string(data), "data failed")
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	res, err := http.Get(ts.URL + "/users") //왜 통과가 될까? => 해당 url이 없으면 / 이 호출됨

	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body) //Body는 Buffer값이라 바로 읽어올 수 없음. 그래서 ioutil을 이용해 버퍼의 내용을 전부 읽어와 data에 저장할 것임
	assert.Equal(string(data), "No Users", "No String")
	//assert.Equal("Hello World", string(data), "data failed")
}

func TestUsers_WithUsersData(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"Lee", "last_name":"hg", "email":"dlgusrb"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode) // stausCreated : 201번

	res, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"Lee2", "last_name":"hg2", "email":"dlgusrb2"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode) // stausCreated : 201번

	res, err = http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	users := []*User{}
	err = json.NewDecoder(res.Body).Decode(&users)
	assert.NoError(err)
	assert.Equal(2, len(users))
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	res, err := http.Get(ts.URL + "/users/98") //왜 통과가 될까? => 해당 url이 없으면 / 이 호출됨

	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body) //Body는 Buffer값이라 바로 읽어올 수 없음. 그래서 ioutil을 이용해 버퍼의 내용을 전부 읽어와 data에 저장할 것임
	assert.Contains(string(data), "No User ID:98", "No String")

}

func TestCreateUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"Lee", "last_name":"hg", "email":"dlgusrb"}`))
	assert.NoError(err)

	assert.Equal(http.StatusCreated, res.StatusCode) // stausCreated : 201번

	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.Equal("Lee", user.FirstName)
	assert.Equal("hg", user.LastName)

	id := user.ID
	res, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(res.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)
}

func TestDeleteUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil) //DELETE 메소드는 지원해주지 않으므로 이와 같이 사용
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User ID") //현재는 Data가 존재하지 않으므로 "No User ID" 가 나와야 함

	res, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"Lee", "last_name":"hg", "email":"dlgusrb"}`)) //Delete할 data 추가
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil) // 추가한 data 삭제
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ = ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "Delete User ID")
}

func TestUpdateUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(`{"id":1,"first_name":"Lack", "last_name":"m", "email":"dlgusrb1234"}`))
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User ID") //현재는 Data가 존재하지 않으므로 "No User ID" 가 나와야 함

	res, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"Lee", "last_name":"hg", "email":"dlgusrb"}`)) //update 할 data 추가
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID) //user가 잘 생성됐는지 확인

	updateStr := fmt.Sprintf(`{"id":%d,"first_name":"Lack"}`, user.ID)
	req, _ = http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(updateStr))
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	updateduser := new(User)
	err = json.NewDecoder(res.Body).Decode(updateduser)
	assert.NoError(err)
	assert.Equal(updateduser.ID, user.ID) //user가 잘 생성됐는지 확인
	assert.Equal(updateduser.FirstName, "Lack")
	assert.Equal(updateduser.LastName, user.LastName)
}
