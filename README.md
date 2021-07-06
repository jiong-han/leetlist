# leetlist

## Overview
Golang program extacts leetcode questions as list with customized filter.

## Feature
* No username, password required
* Customizable filter
* lightweight

## Uasge

```
package main

import "github.com/jionghann/leetlist"

func main() {
	if err := leetlist.Extract("medium.csv", func(q leetlist.Question) bool {
		return q.Difficulty == "Medium"
	}); err != nil {
		panic(err)
	}
}
```

```
export cookie = "{COOKIE_GOES_HERE}"
go run main.go
```

See [example](./example) for more details
