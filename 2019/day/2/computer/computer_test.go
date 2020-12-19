package computer

import (
	"adventofcode/2019/day/2/computer"
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestCompute(t *testing.T) {

	type test struct {
		name  string
		input computer.Program
		want  computer.Program
	}

	tests := []test{
		{
			name:  "(1 + 1 = 2)",
			input: computer.Program{1, 0, 0, 0, 99},
			want:  computer.Program{2, 0, 0, 0, 99},
		},
		{
			name:  "(3 * 2 = 6)",
			input: computer.Program{2, 3, 0, 3, 99},
			want:  computer.Program{2, 3, 0, 6, 99},
		},
		{
			name:  "(99 * 99 = 9801)",
			input: computer.Program{2, 4, 4, 5, 99, 0},
			want:  computer.Program{2, 4, 4, 5, 99, 9801},
		},
		{
			name:  "Complejo",
			input: computer.Program{1, 1, 1, 4, 99, 5, 6, 0, 99},
			want:  computer.Program{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, tc := range tests {
		got := computer.Compute(&tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %[2]T %[2]v, got: %[3]T %[3]v", tc.name, tc.want, got)
		}
	}
}

func TestIsComma(t *testing.T) {

	type test struct {
		input []byte
		want  bool
	}

	tests := []test{
		{
			input: []byte(","),
			want:  true,
		},
		{
			input: []byte("."),
			want:  false,
		},
	}

	for _, tc := range tests {
		r, _ := utf8.DecodeRune(tc.input)
		got := computer.IsComma(r)
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
