package merge

import (
	"sort"
)

func Merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	// Ordering the list of intervals by the first interval value
	// helps us make assumptions that are later needed for the algorithm to work.
	//
	// Why?
	// Ordering by the first item means that for each item the right / next neighbor interval
	// is guaranteed to either start at the same low end or is greater than it.
	// This let's us assume that each right neighbor is either "mergeable" or not,
	// implying that the n-th neighbor is not "mergeable" if the next neighbor is not.
	// Therefore no look-ahead other than a 1-step look-ahead is needed, which enables
	// us to traverse the sorted list linearly.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	res := make([][]int, 0)
	for i := range intervals { // iterate linearly
		// over the list of intervals
		if len(res) == 0 {
			res = append(res, intervals[i])
		}
		currentSoFar := res[len(res)-1]
		if intersects(currentSoFar, intervals[i]) { // if it intersects we create a
			// "merged" interval containing both (the one so far and the current interval we are looking at)
			res[len(res)-1] = mergedInterval(currentSoFar, intervals[i])
		} else { // if not no merge can be performed, just append it and mark it as current
			res = append(res, intervals[i])
		}
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
	return a[1] >= b[0]
}
