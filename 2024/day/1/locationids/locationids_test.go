package locationids_test

import (
	"adventofcode/2024/day/1/locationids"
	"fmt"
	"testing"
)

func TestPopulateVectors(t *testing.T) {
	locs := new(locationids.LocationIDs)
	err := locs.PopulateVectors("../input/test_input.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedLeft := []int{1, 2, 3}
	expectedRight := []int{4, 5, 6}

	if !equal(locs.Left, expectedLeft) {
		t.Errorf("Left vectors are not equal. Expected: %v, Got: %v", expectedLeft, locs.Left)
	}

	if !equal(locs.Right, expectedRight) {
		t.Errorf("Right vectors are not equal. Expected: %v, Got: %v", expectedRight, locs.Right)
	}
}

func TestComputeTotalDistance(t *testing.T) {
	testCases := []struct {
		left             []int
		right            []int
		expectedDistance int
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}, 9},
		{[]int{1, 2}, []int{5, 1}, 5},
		{[]int{1, 2, 3, 3, 3, 4}, []int{3, 3, 3, 4, 5, 9}, 11}, // single element
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestComputeTotalDistance_%d", i), func(t *testing.T) {
			locs := &locationids.LocationIDs{Left: tc.left, Right: tc.right}
			totalDistance := locs.ComputeTotalDistance()
			if totalDistance != tc.expectedDistance {
				t.Errorf("Expected total distance: %d, but got %d", tc.expectedDistance, totalDistance)
			}
		})

	}

}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
