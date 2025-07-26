package main

import (
	"fmt"
	"net/http"

	"github.com/ttetrafon/go-web-server-lib/example/web-api/api"
	"github.com/ttetrafon/go-web-server-lib/middleware"
)

func main() {
	fmt.Println("Starting the example web-api!")

	mux := http.NewServeMux()
	mux.Handle("/", middleware.Chain(http.HandlerFunc(api.GetRoot), middleware.Logging))

}
