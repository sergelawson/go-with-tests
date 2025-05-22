package main

import (
	"fmt"
	"io"
	"os"
)


func Greet(b io.Writer, s string){
	fmt.Fprintf(b, "Hello, %s", s)
}

func main(){
	Greet(os.Stdout, "Serge")
}