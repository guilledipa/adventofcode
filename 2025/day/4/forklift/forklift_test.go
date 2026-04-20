package forklift

import (
	"testing"
)

func TestForklift(t *testing.T) {
	var testCases = []struct {
		in      string
		want    int
		wantErr bool
	}{
		{"../input/test_data.txt", 13, false},
	}
	for _, tc := range testCases {
		got, err := Forklift(tc.in)
		if (err != nil) != tc.wantErr {
			t.Errorf("Forklift(%q) error = %v, wantErr %v", tc.in, err, tc.wantErr)
			continue
		}
		if got != tc.want {
			t.Errorf("Forklift(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
}
