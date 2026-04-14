package dial

import (
	"testing"
)

func TestExecInstruction(t *testing.T) {
	var testCases = []struct {
		in           string
		position     int
		zeros        int
		wantPosition int
		wantZeros    int
		wantErr      bool
	}{
		{"L68", 50, 0, 82, 0, false},
		{"L30", 82, 0, 52, 0, false},
		{"R48", 52, 0, 0, 1, false},
		{"L5", 0, 1, 95, 1, false},
		{"R60", 95, 0, 55, 0, false},
		{"L55", 55, 1, 0, 2, false},
		{"L1", 0, 2, 99, 2, false},
		{"L99", 99, 2, 0, 3, false},
		{"R14", 0, 3, 14, 3, false},
		{"L82", 14, 3, 32, 3, false},
		{"L898", 52, 0, 54, 0, false},
	}
	for _, tc := range testCases {
		dial := Dial{
			Numbers:  dialNumbers,
			Position: tc.position,
			Zeros:    tc.zeros,
		}
		err := dial.execInstruction(tc.in, false)
		if tc.wantErr && err == nil {
			t.Errorf("ExecInstruction(%q) expected error but got none", tc.in)
		}
		if !tc.wantErr && err != nil {
			t.Errorf("ExecInstruction(%q) unexpected error: %v", tc.in, err)
		}
		if dial.Position != tc.wantPosition {
			t.Errorf("ExecInstruction(%q) expected position %d but got %d", tc.in, tc.wantPosition, dial.Position)
		}
		if dial.Zeros != tc.wantZeros {
			t.Errorf("ExecInstruction(%q) expected zeros %d but got %d", tc.in, tc.wantZeros, dial.Zeros)
		}
	}
}

func TestParseInstruction(t *testing.T) {
	var testCases = []struct {
		in        string
		direction string
		clicks    int
		wantErr   bool
	}{
		{"R120", "R", 120, false},
		{"L120", "L", 120, false},
		{"X120", "", 0, true},
	}
	for _, tc := range testCases {
		direction, clicks, err := parseInstruction(tc.in)
		if tc.wantErr && err == nil {
			t.Errorf("parseInstruction(%q) expected error but got none", tc.in)
		}
		if !tc.wantErr && err != nil {
			t.Errorf("parseInstruction(%q) unexpected error: %v", tc.in, err)
		}
		if direction != tc.direction {
			t.Errorf("parseInstruction(%q) expected direction %q but got %q", tc.in, tc.direction, direction)
		}
		if clicks != tc.clicks {
			t.Errorf("parseInstruction(%q) expected clicks %d but got %d", tc.in, tc.clicks, clicks)
		}
	}
}

func TestCountZeroCrossings(t *testing.T) {
	var testCases = []struct {
		direction string
		start     int
		clicks    int
		want      int
	}{
		{"L", 44, 44, 1},
		{"L", 0, 39, 0},
		{"L", 0, 100, 1},
		{"L", 55, 55, 1},
		{"R", 0, 100, 1},
	}
	for _, tc := range testCases {
		got := countZeroCrossings(tc.direction, tc.start, tc.clicks, dialNumbers)
		if got != tc.want {
			t.Errorf("countZeroCrossings(%q, %d, %d) = %d; want %d", tc.direction, tc.start, tc.clicks, got, tc.want)
		}
	}
}
