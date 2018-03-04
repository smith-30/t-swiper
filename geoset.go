package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/disiqueira/tindergo"
)

var (
	concurrent    = 10
	semaphoreChan = make(chan struct{}, concurrent)
)

func main() {
	token := flag.String("token", "", "Your Facebook Token.")

	flag.Parse()

	if *token == "" {
		fmt.Println("You must provide a valid Facebook Token.")
		os.Exit(2)
	}

	t := tindergo.New()

	err := t.Authenticate(*token)
	checkError(err)

	p, err := t.Profile()
}
