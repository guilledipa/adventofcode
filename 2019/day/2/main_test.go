package main

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestCompute(t *testing.T) {

	type test struct {
		name  string
		input Program
		want  Program
	}

	tests := []test{
		{
			name:  "(1 + 1 = 2)",
			input: Program{1, 0, 0, 0, 99},
			want:  Program{2, 0, 0, 0, 99},
		},
		{
			name:  "(3 * 2 = 6)",
			input: Program{2, 3, 0, 3, 99},
			want:  Program{2, 3, 0, 6, 99},
		},
		{
			name:  "(99 * 99 = 9801)",
			input: Program{2, 4, 4, 5, 99, 0},
			want:  Program{2, 4, 4, 5, 99, 9801},
		},
		{
			name:  "Complejo",
			input: Program{1, 1, 1, 4, 99, 5, 6, 0, 99},
			want:  Program{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, tc := range tests {
		got := Compute(&tc.input)
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
		got := isComma(r)
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
