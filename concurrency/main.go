package main

import (
	"fmt"
)

type result struct {
	string
	bool
}

type WebsiteChecker func(url string) bool

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {

	results := make(map[string]bool)
	resultChannel := make(chan result, 100)

	for _, url := range urls {

		go func(u string) {
			resultChannel <- result{u, wc(u)}

		}(url)

	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

func main() {

	urls := []string{"google.com", "archlinux.org", "vk.com"}

	wc := func(url string) bool {
		return false
	}

	res := CheckWebsite(wc, urls)

	fmt.Println(res)

}
