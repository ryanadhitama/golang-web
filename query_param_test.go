package golang_web

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"io"
)

func QueryHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	fmt.Fprintf(writer, "Hello %s", name)
}

func TestHttpQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Ryan", nil)
	recorder := httptest.NewRecorder()

	QueryHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body) 
	fmt.Println(bodyString)
}
