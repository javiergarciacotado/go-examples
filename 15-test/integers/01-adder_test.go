package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	expected := 4
	actual := Add(2, 2)

	if actual != expected {
		t.Errorf("got '%d' - wanted '%d'", actual, expected)
	}
}

func ExampleAdd() {
	actual := Add(1, 5)
	fmt.Println(actual)
	//output: 6

}
