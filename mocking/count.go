package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countDownStart = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct{
	count int
}

type DefaultSleeper struct {
}

func (s *SpySleeper) Sleep(){
	s.count++
}
func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Count(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Count(os.Stdout, &DefaultSleeper{})
}
