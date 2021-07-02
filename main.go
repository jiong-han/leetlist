package main

import (
	"log"
	"os"
	"sync"

	"github.com/jionghann/leetcode-like-dislike/internal"
)

func main() {
	wg := &sync.WaitGroup{}
	cookie := os.Getenv("cookie")
	if len(cookie) == 0 {
		log.Panic("cookie required")
		return
	}

	ch := make(chan internal.QuestionDetail)
	fetcher := internal.NewFetcher(cookie, ch)
	recorder, err := internal.NewRecorder(ch)
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("start recording...")
		if err := recorder.Record(); err != nil {
			log.Println("recorder error: ", err)
		}
	}()

	log.Println("start fetching...")
	if err := fetcher.Fetch(); err != nil {
		log.Println("fetcher error: ", err)
	}

	close(ch)
	wg.Wait()
}
