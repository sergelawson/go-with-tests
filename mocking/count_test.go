package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	buffer := bytes.Buffer{}
	sleeper := SpySleeper{}

	Count(&buffer, &sleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	sleepCount := sleeper.count

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if sleepCount != 3 {
		t.Errorf("Should sleep for 3 seconds but slept for: %d", sleepCount)
	}
}
