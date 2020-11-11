package main

import (
	"reflect"
	"testing"
)

func TestFuelRequirement(t *testing.T) {
	masses := []int{-1, 0, 12, 14, 1969, 100756}
	want := []int{0, 0, 2, 2, 966, 50346}
	got := []int{}
	for _, m := range masses {
		got = append(got, fuelRequirement(m))
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("fuelRequirement(%v) = %v; want %v", masses, got, want)
	}
}
