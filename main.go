package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/disiqueira/tindergo"
	"github.com/k0kubun/pp"
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

	if err != nil {
		return
	}

	fmt.Println("Authentication succeessed !!")

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		semaphoreChan <- struct{}{}
		go func(tinder tindergo.TinderGo) {
			defer func() {
				<-semaphoreChan // read to release a slot
				wg.Done()
			}()

			recs, err := tinder.RecsCore()
			checkError(err)
			if err != nil {
				return
			}

			for _, e := range recs {
				res, err := tinder.Like(e)
				checkError(err)
				pp.Print(res)
			}
		}(*t)
	}

	wg.Wait()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("err... ", err)
	}
}
