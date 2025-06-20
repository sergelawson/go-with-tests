package cntext

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var results strings.Builder

		for _, s := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy got cancelled")
				return

			default:
				time.Sleep(5 * time.Millisecond)
				results.WriteRune(s)
			}

		}
		data <- results.String()

	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case d := <-data:
		return d, nil
	}
}

func TestServer(t *testing.T) {
	t.Run("Should respond with hello, world", func(t *testing.T) {
		data := "Hello, world"
		store := &SpyStore{data, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got : %s, want %s", response.Body.String(), data)
		}

	})

	t.Run("tells store to cancel work if store is canceled", func(t *testing.T) {
		data := "Hello, world is not a good phrase fpr this test"
		store := &SpyStore{data, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())

		time.AfterFunc(2*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}

	})

}
