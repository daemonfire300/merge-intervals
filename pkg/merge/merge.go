package merge

import (
	"sort"
)

func Merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	res := make([][]int, 0)
	for i := range intervals {
		var current []int
		if len(res) > 0 {
			current = res[len(res)-1] // last added interval
		}

		if current != nil {
			// check if the next entry (i) can be merged with current
			if intersects(intervals[i], current) {
				// if so, update current
				res[len(res)-1] = mergedInterval(intervals[i], current)
				continue
			}
		}

		// the current interval does not intersect with the next entry (i), let's find overlaps in the upcoming
		// entries (j)

		// compare i to every other entry (j) and select largest interval
		currentSoFar := intervals[i]
		for j := range intervals[i:] {
			if intersects(currentSoFar, intervals[j]) {
				currentSoFar = mergedInterval(currentSoFar, intervals[j])
			}
		}
		res = append(res, currentSoFar)
	}
	return res
}

// mergedInterval merges two intervals and returns an interval that captures both.
//
// Example:
// in: [1, 3], [3, 6]
// out: [1, 6]
func mergedInterval(a, b []int) []int {
	res := make([]int, 2)
	if a[0] < b[0] {
		res[0] = a[0]
	} else {
		res[0] = b[0]
	}

	if a[1] > b[1] {
		res[1] = a[1]
	} else {
		res[1] = b[1]
	}
	return res
}

// intersects, answers: a fits b?
func intersects(a, b []int) bool {
	return ((a[0] >= b[0] && a[0] <= b[1]) || (a[1] >= b[1] && a[1] <= b[0])) ||
		(a[0] <= b[0] && a[1] >= b[1]) || (a[1] == b[0])
}
