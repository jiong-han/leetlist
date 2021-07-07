package main

import "github.com/jionghann/leetlist"

func main() {
	if err := leetlist.Extract("medium.csv", func(q leetlist.Question) bool {
		return (q.Difficulty == "Medium" || q.Difficulty == "Easy") && q.Likes > q.Dislikes && q.Freq > 50.0
	}); err != nil {
		panic(err)
	}
}
