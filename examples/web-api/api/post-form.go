package api

import (
	"fmt"
	"io"
	"net/http"
)

func PostForm(w http.ResponseWriter, r *http.Request) {
	myName := r.PostFormValue("name")
	if myName == "" {
		myName = "HTTP"
	}
	io.WriteString(w, fmt.Sprintf("Hello, %s!\n", myName))
}
