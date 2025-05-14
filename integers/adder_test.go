package integers

import (
	"fmt"
	"testing"
)


func TestAdder(t *testing.T){
	want := 4
	got := Add(2, 2)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}


func ExampleAdd(){
	sum :=  Add(2, 1)
	fmt.Println(sum)
	// Output: 3	
}