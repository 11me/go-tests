package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "example.com" {
		return false
	}
	return true
}

func slowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return false
}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsite(slowWebsiteChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T) {

	urls := []string{
		"google.com",
		"vk.com",
		"yahoo.com",
		"duckduckgo.com",
		"github.com",
		"example.com",
	}

	got := CheckWebsite(mockWebsiteChecker, urls)
	want := map[string]bool{
		"google.com":     true,
		"vk.com":         true,
		"yahoo.com":      true,
		"duckduckgo.com": true,
		"github.com":     true,
		"example.com":    false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}
