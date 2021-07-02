package internal

import (
	"encoding/csv"
	"log"
	"os"
)

type Recorder struct {
	writer *csv.Writer
	ch     chan QuestionDetail
}

func NewRecorder(ch chan QuestionDetail) (Recorder, error) {
	file, err := os.Create("./list.csv")
	if err != nil {
		return Recorder{}, err
	}

	writer := csv.NewWriter(file)
	err = writer.Write([]string{"#", "title", "acRate", "difficulty", "freq", "likes", "dislikes", "link"})
	if err != nil {
		return Recorder{}, err
	}

	return Recorder{
		writer: writer,
		ch:     ch,
	}, nil
}

func (r *Recorder) Record() error {
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
