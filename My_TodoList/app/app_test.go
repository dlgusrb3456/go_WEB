package app

import (
	"encoding/json"
	"fmt"
	"go_WEB/My_TodoList/model"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "?sr06468sr"
	DB_NAME     = "todo_1"
)

func TestTodos(t *testing.T) {
	//로그인이 된 후에 다음 테스트 코드가 수행되므로 로그인을 수행해야함.
	//로그인이 됐는지 여부를 확인하는 app.go의 getSessionId()를 func()포인터를 가지는 variable로 만들어서 문제를 해결한다.
	getSessionID = func(r *http.Request) string {
		return "testsessionId"
	}
	// function variable이기 때문에 값을 변경할 수 있음
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	os.Remove("./test.db")
	assert := assert.New(t)
	ah := NewRouter(dbinfo)
	defer ah.Close()

	ts := httptest.NewServer(ah)
	defer ts.Close()

	resp, err := http.PostForm(ts.URL+"/TodoList", url.Values{"name": {"take a walk"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	var todo model.Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "take a walk")
	id1 := todo.ID

	resp, err = http.PostForm(ts.URL+"/TodoList", url.Values{"name": {"take a walk2"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "take a walk2")
	id2 := todo.ID

	resp, err = http.Get(ts.URL + "/TodoList")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos := []*model.Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 2)

	for _, t := range todos {
		if t.ID == id1 {
			assert.Equal("take a walk", t.Name)
		} else if t.ID == id2 {
			assert.Equal("take a walk2", t.Name)
		} else {
			assert.Error(fmt.Errorf("No id in todos"))
		}
	}

	resp, err = http.Get(ts.URL + "/complete-todo/" + strconv.Itoa(id1))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/TodoList")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 2)

	for _, t := range todos {
		if t.ID == id1 {
			assert.True(t.Completed)
		}
	}

	req, _ := http.NewRequest("DELETE", ts.URL+"/TodoList/"+strconv.Itoa(id1), nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/TodoList")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 1)

	for _, t := range todos {
		assert.Equal(t.ID, id2)
	}

	resp, err = http.Get(ts.URL + "/getInfoList/" + strconv.Itoa(id2))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "take a walk2")
}
