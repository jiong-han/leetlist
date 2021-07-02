package internal

import "fmt"

// type questionResponse struct {
// 	Data struct {
// 		Question struct {
// 			Likes    int `json:"likes"`
// 			Dislikes int `json:"dislikes"`
// 		} `json:"question"`
// 	} `json:"data"`
// }

type problemListResponse struct {
	Data struct {
		ProblemSet struct {
			Questions []QuestionDetail `json:"questions"`
			Total     int              `json:"total"`
		} `json:"problemsetQuestionList"`
	} `json:"data"`
}

type QuestionDetail struct {
	AcRate     float64 `json:"acRate"`
	Difficulty string  `json:"difficulty"`
	Freq       float64 `json:"freqBar"`
	Title      string  `json:"title"`
	TitleSlug  string  `json:"titleSlug"`
	Likes      int     `json:"likes"`
	Dislikes   int     `json:"dislikes"`
}

func (qd QuestionDetail) AsStringArr() []string {
	arr := []string{
		qd.Title,
		fmt.Sprintf("%f", qd.AcRate),
		qd.Difficulty,
		fmt.Sprintf("%f", qd.Freq),
		fmt.Sprintf("%d", qd.Likes),
		fmt.Sprintf("%d", qd.Dislikes),
		fmt.Sprintf("https://leetcode.com/problems/%s/", qd.TitleSlug),
	}

	return arr
}
