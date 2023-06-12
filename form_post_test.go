package golang_web

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
	"strings"
	"io"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "%s %s", firstName, lastName)
}

func TestForm(t *testing.T) {
	reqBody := strings.NewReader("first_name=ryan&last_name=adhitama")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080?name=Ryan", reqBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body) 
	fmt.Println(bodyString)
}