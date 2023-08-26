package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := adder(2, 2)
	want := 4

	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
