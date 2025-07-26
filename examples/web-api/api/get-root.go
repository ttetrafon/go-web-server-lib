package api

import (
	"io"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Server is running!\n")
}
