package leetlist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	graphQLEndpoint = "https://leetcode.com/graphql"
	problemSetGql   = "{\"query\":\"\\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\\n  problemsetQuestionList: questionList(\\n    categorySlug: $categorySlug\\n    limit: $limit\\n    skip: $skip\\n    filters: $filters\\n  ) {\\n    total: totalNum\\n    questions: data {\\n      likes\\n      dislikes\\n      acRate\\n      difficulty\\n      freqBar\\n      frontendQuestionId: questionFrontendId\\n      isFavor\\n      paidOnly: isPaidOnly\\n      status\\n      title\\n      titleSlug\\n      topicTags {\\n        name\\n        id\\n        slug\\n      }\\n      hasSolution\\n      hasVideoSolution\\n    }\\n  }\\n}\\n    \",\"variables\":{\"categorySlug\":\"\",\"skip\":%d,\"limit\":%d,\"filters\":{}}}"
)

type fetcher struct {
	cookie string
	client http.Client
	ch     chan Question
	filter Filter
}

func NewFetcher(cookie string, ch chan Question, filter Filter) fetcher {
	return fetcher{
		cookie: cookie,
		client: http.Client{
			Timeout: 3 * time.Second,
		},
		ch:     ch,
		filter: filter,
	}
}

func (f *fetcher) fetch() error {
	skip := 0
	limit := 200
	interval := time.Second
	lastRequestTime := time.Unix(0, 0)

	var resp problemListResponse

	for {
		log.Printf("placing request for range %d - %d", skip, skip+limit)
		for time.Since(lastRequestTime) < interval {
			time.Sleep(interval)
		}

		data, err := f.sendRequest(graphQLEndpoint, fmt.Sprintf(problemSetGql, skip, limit))
		if err != nil {
			log.Println("error sending request: ", err)
			return err
		}

		log.Println("marshalling responses...")
		err = json.Unmarshal(data, &resp)
		if err != nil {
			log.Println("error marshal response: ", err)
			return err
		}

		for _, q := range resp.Data.ProblemSet.Questions {
			log.Printf("checking questions: %s", q.Title)
			if f.filter(q) {
				f.ch <- q
			}
		}

		if len(resp.Data.ProblemSet.Questions) < limit {
			log.Println("No more questions, break")
			return nil
		}

		skip += limit

	}

}

func (f *fetcher) sendRequest(url, reqBody string) ([]byte, error) {
	requestBody := bytes.NewBuffer([]byte(reqBody))
	req, err := http.NewRequest(http.MethodPost, url, requestBody)
	if err != nil {
		log.Println("error creating new request: ", err)
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", f.cookie)
	req.Header.Set("referer", "https://leetcode.com/problemset/all/")
	req.Header.Set("origin", "https://leetcode.com")

	for _, cookie := range req.Cookies() {
		if cookie.Name == "csrftoken" {
			req.Header.Set("x-csrftoken", cookie.Value)
		}
	}

	resp, err := f.client.Do(req)
	if err != nil {
		log.Println("error placing request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error read response body: ", err)
		return nil, err
	}

	return respData, nil

}
