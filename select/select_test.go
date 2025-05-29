package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("should return the fastest website url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL

		got, error := WebsiteRacer(slowURL, want)

		if error != nil {
			t.Errorf("returned an unexpected error %s", error.Error())
		}

		if got != want {
			t.Errorf("want: %q; got: %q", want, got)
		}
	})

	t.Run("should return timeout error website", func(t *testing.T) {
		slowServer := makeDelayedServer(11 * time.Second)
		fastServer := makeDelayedServer(12 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL

		_, error := WebsiteRacer(slowURL, want)

		if error != ErrTimeout {
			t.Errorf("should have return %s", ErrTimeout.Error())
		}

	})
}

func makeDelayedServer(duration time.Duration) *httptest.Server {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)

		w.WriteHeader(http.StatusOK)
	}))

	return server

}
