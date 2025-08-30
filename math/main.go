package main

import (
	"os"
	"time"

	"github.com/sergelawson/go-with-tests/math/clockface"
)

func main() {
	t := time.Now()

	clockface.SVGWriter(os.Stdout, t)
}
