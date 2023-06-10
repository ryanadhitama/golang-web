package golang_web

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
)

func HeaderHandler(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	writer.Header().Add("content-type", "application/json")
	fmt.Fprintf(writer, contentType)
}

func TestHttpHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Ryan", nil)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	HeaderHandler(recorder, request)

	response := recorder.Result()
	fmt.Println(response.Header.Get("content-type"))
}
