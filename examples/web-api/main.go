package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/ttetrafon/go-web-server-lib/example/web-api/api"
	"github.com/ttetrafon/go-web-server-lib/middleware"
)

type ServerAddress string

const keyServerAddr ServerAddress = "serverAddr"

func main() {
	fmt.Println("Starting the example web-api!")

	mux := http.NewServeMux()
	mux.Handle("/", middleware.Chain(http.HandlerFunc(api.GetRoot), middleware.Logging))
	mux.Handle("/hello", middleware.Chain(http.HandlerFunc(api.GetHello), middleware.Logging))
	mux.Handle("/search", middleware.Chain(http.HandlerFunc(api.GetSearch), middleware.Logging))
	mux.Handle("/postData", middleware.Chain(http.HandlerFunc(api.PostData), middleware.Logging, middleware.Authenticate))
	mux.Handle("/helloForm", middleware.Chain(http.HandlerFunc(api.PostForm), middleware.Logging, middleware.Authenticate))
	mux.Handle("/respondWithCode", middleware.Chain(http.HandlerFunc(api.RespondWithCode), middleware.Logging))

	ctx, cancelCtx := context.WithCancel(context.Background())

	fmt.Println("... creating server at :3333!")
	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	fmt.Println("... creating server at :4444!")
	serverTwo := &http.Server{
		Addr:    ":4444",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()

	go func() {
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server two closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server two: %s\n", err)
		}
		cancelCtx()
	}()

	<-ctx.Done()
}
