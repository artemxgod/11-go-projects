package main

import (
	"log"

	"github.com/artemxgod/11-go-projects/email-verifier/pkg/checker"
)

func main() {
	if err := checker.StartCheck(); err != nil {
		log.Fatal(err)
	}
}