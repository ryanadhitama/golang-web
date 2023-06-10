package golang_web

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"io"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body) 
	fmt.Println(bodyString)
}
