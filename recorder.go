package leetlist

import (
	"encoding/csv"
	"log"
	"os"
)

type recorder struct {
	writer *csv.Writer
	ch     chan Question
}

func NewRecorder(filePath string, ch chan Question) (recorder, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return recorder{}, err
	}

	writer := csv.NewWriter(file)
	err = writer.Write([]string{"#", "title", "acRate", "difficulty", "freq", "likes", "dislikes", "link"})
	if err != nil {
		return recorder{}, err
	}

	return recorder{
		writer: writer,
		ch:     ch,
	}, nil
}

func (r *recorder) record() error {
	for qd := range r.ch {

		log.Printf("Recorder processing: %s", qd.Title)

		err := r.writer.Write(qd.AsStringArr())
		if err != nil {
			log.Println("error record: ", err)
			continue
		}
	}

	r.writer.Flush()

	return nil
}
