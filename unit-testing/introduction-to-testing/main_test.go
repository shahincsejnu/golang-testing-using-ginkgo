// run the test using `go test`
// run the test using `go test -v`

package main

import "testing"

func TestSum(t *testing.T) {
	if Sum(10, 20) != 30 {
		t.Errorf("Wrong addition")
	}
}

type Input struct {
	x      int
	y      int
	result int
}

func TestFable(t *testing.T) {
	test := []Input{
		{10, 20, 30},
		{20, 30, 50},
		{1, 1, 2},
		{2, 2, 4},
	}

	for _, values := range test {
		if val := Sum(values.x, values.y); val != values.result {
			t.Errorf("expected: %v, found: %v", values.result, val)
		}
	}
}

func TestTable(t *testing.T) {
	test := []struct {
		x      int
		y      int
		result int
	}{
		{1, 2, 3},
		{2, 2, 5},
	}

	for _, values := range test {
		if val := Sum(values.x, values.y); val != values.result {
			t.Errorf("expected: %v, found: %v", values.result, val)
		}
	}
}

type Demo struct {
	Another *Temp
	Name    *string
}

type Temp struct {
	Name    *string
	Address *string
}
