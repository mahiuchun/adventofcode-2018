package game

import "testing"

func TestHighScore(t *testing.T) {
	var tests = []struct {
		np   int
		nm   int
		want int
	}{
		{9, 25, 32},
		{10, 1618, 8317},
		{13, 7999, 146373},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
	}
	for _, test := range tests {
		if got := HighScore(test.np, test.nm); got != test.want {
			t.Errorf("HighScore(%v, %v) = %v", test.np, test.nm, got)
		}
	}
}
