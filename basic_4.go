package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Food{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
			if f, ok := queryCache(id, m); ok {
				fmt.Println("from cache")
				fmt.Println(f)
			}
			wg.Done()
		}(id, wg, m)
		go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
			if f, ok := queryDatabase(id, m); ok {
				fmt.Println("from database")
				fmt.Println(f)

			}
			wg.Done()
		}(id, wg, m)

		fmt.Printf("Food eith id: '%v' not found", id)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryCache(id int, m *sync.Mutex) (Food, bool) {
	m.Lock()
	f, ok := cache[id]
	m.Unlock()
	return f, ok
}

func queryDatabase(id int, m *sync.Mutex) (Food, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, f := range foods {
		if f.ID == id {
			m.Lock()
			cache[id] = f
			m.Unlock()
			return f, true
		}
	}

	return Food{}, false
}
