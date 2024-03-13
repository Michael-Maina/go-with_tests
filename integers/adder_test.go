package integers

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T) {
	sum := Add(3, 3)
	expected := 6

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}
