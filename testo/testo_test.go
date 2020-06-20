package testo

import "testing"

type testpair struct {
	values  []float64
	average float64
	sum     float64
}

var tests = []testpair{
	{[]float64{1, 3}, 1.5, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 1, 6},
	{[]float64{-1, 1}, 0, 0},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestSummarized(t *testing.T) {
	for _, pair := range tests {
		v := Summarized(pair.values)
		if v != pair.sum {
			t.Error(
				"For", pair.values,
				"expected", pair.sum,
				"got", v,
			)
		}
	}
}
