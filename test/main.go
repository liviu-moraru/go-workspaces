package main

import (
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	var eg errgroup.Group

	eg.Go(func() error {
		log.Println("did a thing")
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Println("err")
	}
}
