package main

import (
	"context"
	"log"

	kiteinstruments "github.com/nsvirk/gokiteinstruments"
)

func main() {
	ctx := context.Background()
	client, err := kiteinstruments.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	runQueries(client)
}
