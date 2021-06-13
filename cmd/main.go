package main

import (
	"log"

	"github.com/codershangfeng/vote-service/app/internal/context"
)

func main() {
	api, err := context.NewAPIHandler(nil)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.AppContext{Port: 8080}

	server := ctx.NewServer(api)
	defer server.Shutdown()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
