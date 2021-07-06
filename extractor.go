package leetlist

import (
	"errors"
	"log"
	"os"
	"sync"
)

func Extract(filePath string, filterFunc Filter) error {
	wg := &sync.WaitGroup{}
	cookie := os.Getenv("cookie")
	if len(cookie) == 0 {
		return errors.New("error cookie noshow")
	}

	ch := make(chan Question)
	fetcher := NewFetcher(cookie, ch, filterFunc)
	recorder, err := NewRecorder(filePath, ch)
	if err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("start recording...")
		if err := recorder.record(); err != nil {
			log.Println("recorder error: ", err)
		}
	}()

	log.Println("start fetching...")
	if err := fetcher.fetch(); err != nil {
		log.Println("fetcher error: ", err)
	}

	close(ch)
	wg.Wait()

	return nil
}
