package size

import (
	"testing"
)

type Test struct {
	in int
	out string
}

var tests = []Test{
	{-1, "nagative"},
	{5, "small"},
}

func TestSize(t *testing.T)  {
	for i , test :=  range test{
		size := Size(test.in)
		if(size != test.out){
			t.Error("#d: Size(%d)")
		}
	}
}