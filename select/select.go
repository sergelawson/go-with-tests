package main

import (
	"errors"
	"net/http"
	"time"
)

var ErrTimeout  = errors.New("timeout, request delayed over 10 seconds")

func WebsiteRacer(a, b string) (string, error) {

	select {
	case <- ping(a): 

	return a, nil

	case <- ping(b):
		return b, nil

	case <- time.After(10 * time.Second): 

	return "", ErrTimeout
	}
}


func ping(url string) (chan struct{}){

	channel := make(chan struct{})

	go func(){

		http.Get(url)
		close(channel)

	}()

	return channel
}

func requestDuration(url string) time.Duration {
	startA := time.Now()
	http.Get(url)
	durationA := time.Since(startA)

	return durationA
}