package tachyonmanifold

import "testing"

func TestCalculateTotal(t *testing.T) {
	tc := []struct {
		file      string
		wantPart1 int
		wantPart2 int
		wantErr   bool
	}{
		{"../input/test_data.txt", 21, 0, false},
	}

	for _, c := range tc {
		gotPart1, gotPart2, err := CalculateTotal(c.file)
		if (err != nil) != c.wantErr {
			t.Errorf("CalculateTotal(%q) error = %v, wantErr %v", c.file, err, c.wantErr)
			continue
		}
		if gotPart1 != c.wantPart1 || gotPart2 != c.wantPart2 {
			t.Errorf("CalculateTotal(%q) = (%v, %v), want (%v, %v)", c.file, gotPart1, gotPart2, c.wantPart1, c.wantPart2)
		}
	}
}
