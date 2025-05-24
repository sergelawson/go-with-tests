package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCount(t *testing.T) {

	t.Run("test sleep and message", func(t *testing.T) {
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

	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeper := SpyCountDownOperations{}

		Count(&spySleeper, &spySleeper)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeper.Calls) {
			t.Errorf("want %v, got %v", want, spySleeper.Calls)
		}

	})

}


func TestConfigurableSleeper(t *testing.T) {
	duration := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{duration: duration, sleep: spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != duration {
		t.Errorf("should have slept for %v but slept for %v", duration, spyTime.durationSlept)

	}
}