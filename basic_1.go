package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Food{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		if f, ok := queryCache(id); ok {
			fmt.Println("from cache")
			fmt.Println(f)
			continue
		}
		if f, ok := queryDatabase(id); ok {
			fmt.Println("from database")
			fmt.Println(f)
			continue
		}

		fmt.Printf("Food eith id: '%v' not found", id)
		time.Sleep(150 * time.Millisecond)
	}
}

func queryCache(id int) (Food, bool) {
	f, ok := cache[id]
	return f, ok
}

func queryDatabase(id int) (Food, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, f := range foods {
		if f.ID == id {
			cache[id] = f
			return f, true
		}
	}

	return Food{}, false
}
