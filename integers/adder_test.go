package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	sum := Add(1, 2)
	expected := 3

	if sum != expected {
		t.Errorf("got %d, want %d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
