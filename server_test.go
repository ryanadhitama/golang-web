package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// Test Server
func TestServer(t *testing.T) {
	server := http.Server {
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

// Test Handler
func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	}
	server := http.Server {
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

// ServeMux
func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	mux.HandleFunc("/ryan", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello Ryan")
	})

	server := http.Server {
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

// Request
func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprint(writer, request.RequestURI)
	}
	server := http.Server {
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}