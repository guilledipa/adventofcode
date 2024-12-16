package locationids

import (
	"fmt"
	"testing"
)

func TestPopulateVectors(t *testing.T) {
	locs := new(LocationIDs)
	err := locs.populateVectors("test_input.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedLeft := []int{1, 2, 3}
	expectedRight := []int{4, 5, 6}

	if !equal(locs.left, expectedLeft) {
		t.Errorf("Left vectors are not equal. Expected: %v, Got: %v", expectedLeft, locs.left)
	}

	if !equal(locs.right, expectedRight) {
		t.Errorf("Right vectors are not equal. Expected: %v, Got: %v", expectedRight, locs.right)
	}
}

func TestComputeTotalDistance(t *testing.T) {
	testCases := []struct {
		left             []int
		right            []int
		expectedDistance float64
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}, float64(len([]int{1, 2, 3}) * 3)}, // simple case
		{[]int{1, 2}, []int{5, 1}, 5},
		{[]int{1}, []int{1}, 0}, // single element
		{[]int{}, []int{}, 0},   // empty vectors
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestComputeTotalDistance_%d", i), func(t *testing.T) {
			locs := &LocationIDs{left: tc.left, right: tc.right}
			totalDistance := locs.computeTotalDistance()
			if totalDistance != tc.expectedDistance {
				t.Errorf("Expected total distance: %f, but got %f", tc.expectedDistance, totalDistance)
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
