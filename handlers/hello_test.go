package handlers

import "testing"

func TestSomeNumber(t *testing.T) {
	positiveResult := SomeNumber(10)

	if positiveResult != 52 {
		t.Errorf("SomeNumber(10) = %d; want 52", positiveResult)
	}
}
