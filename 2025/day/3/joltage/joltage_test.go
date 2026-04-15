package joltage

import (
	"testing"
)

func TestMaxJoltage(t *testing.T) {
	var testCases = []struct {
		in      string
		want    int
		wantErr bool
	}{
		{"../input/test_data.txt", 357, false},
	}
	for _, tc := range testCases {
		got, err := MaxJoltage(tc.in)
		if (err != nil) != tc.wantErr {
			t.Errorf("MaxJoltage(%q) error = %v, wantErr %v", tc.in, err, tc.wantErr)
			continue
		}
		if got != tc.want {
			t.Errorf("MaxJoltage(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
}
