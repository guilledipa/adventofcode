package fresh

import (
	"testing"
)

func TestChecker(t *testing.T) {
	tc := []struct {
		file      string
		wantPart1 int
		wantPart2 int
		wantErr   bool
	}{
		{"../input/test_data.txt", 3, 0, false},
	}

	for _, c := range tc {
		gotPart1, gotPart2, err := Checker(c.file)
		if (err != nil) != c.wantErr {
			t.Errorf("Checker(%q) error = %v, wantErr %v", c.file, err, c.wantErr)
			continue
		}
		if gotPart1 != c.wantPart1 || gotPart2 != c.wantPart2 {
			t.Errorf("Checker(%q) = (%v, %v), want (%v, %v)", c.file, gotPart1, gotPart2, c.wantPart1, c.wantPart2)
		}
	}
}
