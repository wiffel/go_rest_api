package main

import (
	"log"
	"flag"

	"github.com/wiff-s/go_rest/internal/apiserver"
)

var (
	configPath string
)

func init() {
	flag.
}

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
