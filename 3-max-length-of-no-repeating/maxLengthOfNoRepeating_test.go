package main

import "testing"

func Test_maxLengthOfNoRepeating(t *testing.T) {
	dataMap := []struct{
		str string
		expected int
	} {
		{"aabbccdef", 4},
		{"灰化肥挥发发灰会发黑", 5},
	}

	for _, v := range dataMap {
		maxLen := maxLengthOfNoRepeating(v.str)
		if maxLen != v.expected {
			t.Errorf("max length of \"%s\" expected %d. Got %d",
				v.str, v.expected, maxLen)
		}
	}
}

func Benchmark_maxLengthOfNoRepeating(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxLengthOfNoRepeating("黑化肥挥发发灰会发货灰度")
	}
}
