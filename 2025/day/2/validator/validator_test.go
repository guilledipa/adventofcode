package validator

import (
	"testing"
)

func TestValidate(t *testing.T) {

	var testCases = []struct {
		in      string
		wantP1  int
		wantP2  int
		wantErr bool
	}{
		{"../input/test_data.txt", 1227775554, 4174379265, false},
	}
	for _, tc := range testCases {
		gotP1, gotP2, err := Validate(tc.in)
		if (err != nil) != tc.wantErr {
			t.Errorf("Validate(%q) error = %v, wantErr %v", tc.in, err, tc.wantErr)
			continue
		}
		if gotP1 != tc.wantP1 || gotP2 != tc.wantP2 {
			t.Errorf("Validate(%q) = (%v, %v), want (%v, %v)", tc.in, gotP1, gotP2, tc.wantP1, tc.wantP2)
		}
	}
}
