package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestWebsiteChecker(t *testing.T){

	urls := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	got := CheckWebsites(mockWebsiteChecker, urls)

	want := map[string]bool{
		"http://google.com": true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds": false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
	
}

func slowWebsiteChecker (_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i, _ := range urls {
		urls[i] = "Test test"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}