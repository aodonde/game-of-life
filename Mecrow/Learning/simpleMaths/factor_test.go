package simpleMaths

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsPrime(t *testing.T) {
	cases := []struct {
		input          int
		correct_output bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{51, false},
		{17, true},
		{97, true},
		{121, false},
	}
	for _, c := range cases {
		received_output := IsPrime(c.input)
		if received_output != c.correct_output {
			t.Errorf("IsPrime(%d) == %t, should be %t",
				c.input, received_output, c.correct_output)
		}
	}
}

func TestGetNextPrime(t *testing.T) {
	cases := []struct {
		input          int
		correct_output int
	}{
		{1, 2},
		{2, 3},
		{3, 5},
		{4, 5},
		{51, 53},
		{17, 19},
		{97, 101},
		{121, 127},
	}
	for _, c := range cases {
		received_output := getNextPrime(c.input)
		if received_output != c.correct_output {
			t.Errorf("IsPrime(%d) == %d, should be %d",
				c.input, received_output, c.correct_output)
		}
	}
}

func TestPrimeFactors(t *testing.T) {
	cases := []struct {
		input          int
		correct_output []int
	}{
		{1, []int{}},
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2, 2}},
		{51, []int{3, 17}},
		{17, []int{17}},
		{97, []int{97}},
		{121, []int{11, 11}},
	}
	for _, c := range cases {
		received_output := PrimeFactors(c.input)
		if !cmp.Equal(received_output, c.correct_output) {
			t.Errorf("IsPrime(%d) == %v, should be %v",
				c.input, received_output, c.correct_output)
		}
	}
}
