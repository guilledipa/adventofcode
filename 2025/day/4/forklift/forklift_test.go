package forklift

import (
	"testing"
)

func TestForklift(t *testing.T) {
	var testCases = []struct {
		in      string
		want    int
		want2   int
		wantErr bool
	}{
		{"../input/test_data.txt", 13, 43, false},
	}
	for _, tc := range testCases {
		got, got2, err := Forklift(tc.in)
		if (err != nil) != tc.wantErr {
			t.Errorf("Forklift(%q) error = %v, wantErr %v", tc.in, err, tc.wantErr)
			continue
		}
		if got != tc.want || got2 != tc.want2 {
			t.Errorf("Forklift(%q) = (%v, %v), want (%v, %v)", tc.in, got, got2, tc.want, tc.want2)
		}
	}
}
