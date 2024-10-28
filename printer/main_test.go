package printer

import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{5, "short"},
	{0,"zero"},
	{55,"long"},
	{130,"very long"},
}

func TestLength(t *testing.T) {
	for i, test := range tests {
		size := Length(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}