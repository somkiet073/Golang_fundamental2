package fizzbuzz_test

import (
	"hellogo/fizzbuzz"
	"testing"
)

func TestInput1ShowBeDisplay1(t *testing.T) {
	v := fizzbuzz.Fizzbuzz(1)
	if "1" != v {
		t.Error("fizzbuzz of 1 should be '1' but have", v)
	}
}
