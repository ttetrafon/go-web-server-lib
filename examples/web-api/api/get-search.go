package api

import (
	"fmt"
	"io"
	"net/http"
)

func GetSearch(w http.ResponseWriter, r *http.Request) {
	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")
	hasFilter := r.URL.Query().Has("filter")
	filter := r.URL.Query().Get("filter")

	fmt.Printf("request.params: first(%t)=%s, second(%t)=%s, filter(%t)=%s\n",
		hasFirst, first,
		hasSecond, second,
		hasFilter, filter,
	)
	io.WriteString(w, "Search completed!\n")
}
