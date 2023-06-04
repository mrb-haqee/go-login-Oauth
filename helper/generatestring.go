package helper

import (
	"math/rand"
	"strings"
	"sync"
)

func GenerateRandomString() (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result strings.Builder
	var wg sync.WaitGroup
	ch := make(chan string)

	makeStr := func(ch chan string) {
		defer wg.Done()
		var str strings.Builder
		for i := 0; i < 20; i++ {
			str.WriteString(string(letters[rand.Intn(len(letters))]))
		}
		ch <- str.String()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go makeStr(ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for e := range ch {
		result.WriteString(e)
	}

	return result.String(), nil
}