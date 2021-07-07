package leetlist

import (
	"fmt"
	"strconv"
)

type Filter func(Question) bool

type Question struct {
	ID         string  `json:"frontendQuestionId"`
	AcRate     float64 `json:"acRate"`
	Difficulty string  `json:"difficulty"`
	Freq       float64 `json:"freqBar"`
	Title      string  `json:"title"`
	TitleSlug  string  `json:"titleSlug"`
	Likes      int     `json:"likes"`
	Dislikes   int     `json:"dislikes"`
}

func (q Question) AsStringArr() []string {
	arr := []string{
		q.ID,
		q.Title,
		fmt.Sprintf("%f", q.AcRate),
		q.Difficulty,
		fmt.Sprintf("%f", q.Freq),
		strconv.Itoa(q.Likes),
		strconv.Itoa(q.Dislikes),
		fmt.Sprintf("https://leetcode.com/problems/%s/", q.TitleSlug),
	}

	return arr
}

type problemListResponse struct {
	Data struct {
		ProblemSet struct {
			Questions []Question `json:"questions"`
			Total     int        `json:"total"`
		} `json:"problemsetQuestionList"`
	} `json:"data"`
}
