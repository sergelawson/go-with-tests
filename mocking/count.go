package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countDownStart = 3
const finalWord = "Go!"
const write = "write"
const sleep = "sleep"



type SpyCountDownOperations struct{
	Calls []string
}

type Sleeper interface {
	Sleep()
}

type SpyTime struct {
	durationSlept time.Duration
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

type SpySleeper struct{
	count int
}

type DefaultSleeper struct {
}

func (s *SpyTime) Sleep(duration time.Duration){
	s.durationSlept = duration
}

func (s *SpySleeper) Sleep(){
	s.count++
}
func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpyCountDownOperations) Sleep(){
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(b []byte) (n int, er error){
	s.Calls = append(s.Calls, write)
	return
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Count(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := ConfigurableSleeper{1*time.Second, time.Sleep}
	Count(os.Stdout, &sleeper)
}
