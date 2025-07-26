package api

import (
	"io"
	"net/http"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, HTTP!\n")
}
