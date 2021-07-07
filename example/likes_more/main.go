package main

import "github.com/jionghann/leetlist"

func main() {
	if err := leetlist.Extract("likes.csv", func(q leetlist.Question) bool {
		return q.Likes > q.Dislikes
	}); err != nil {
		panic(err)
	}
}
