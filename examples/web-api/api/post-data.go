package api

import (
	"fmt"
	"io"
	"net/http"
)

func PostData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}

	fmt.Printf("request.body:\n%s\n", body)
	io.WriteString(w, "Data posted...!\n")
}
