package binary_search

import (
	"testing"

	"github.com/franela/goblin"
)

// taken from https://stackoverflow.com/a/39868255/7216271
func makeRangedList(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// taken from https://stackoverflow.com/a/37335777/7216271
func removeElementFromList(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func TestBinarySearch(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Binary Search Unit Tests", func() {
		g.It("Look for every value in list from 1 to 1000, without mutating list", func() {
			sortedList := makeRangedList(1, 1000)
			sortedListWeStartedWith := makeCopy(sortedList)
			for idx, val := range sortedList {
				result := findValueInList(val, sortedList)
				g.Assert(result).Equal(idx)
			}
			g.Assert(sortedList).Equal(sortedListWeStartedWith)
		})

		g.It("Look for 3000 in empty list", func() {
			sortedList := []int{}
			result := findValueInList(3000, sortedList)
			g.Assert(result).Equal(-1)
		})

		g.It("Look for 3000 in list with single value 4", func() {
			sortedList := []int{4}
			result := findValueInList(3000, sortedList)
			g.Assert(result).Equal(-1)
		})

		g.It("Look for 3000 in list from 1 to 5", func() {
			sortedList := makeRangedList(1, 5)
			result := findValueInList(3000, sortedList)
			g.Assert(result).Equal(-1)
		})

		g.It("Look for 3000 in list with single value 3000", func() {
			sortedList := []int{3000}
			result := findValueInList(3000, sortedList)
			g.Assert(result).Equal(0)
		})

		g.It("Look for 3000 in list from 1 to 5487", func() {
			sortedList := makeRangedList(1, 5487)
			result := findValueInList(3000, sortedList)
			g.Assert(result).Equal(2999)
		})

		g.It("Look for 3000 in list from 1 to 5487 with missing value 3000", func() {
			sortedList := removeElementFromList(makeRangedList(1, 5487), 2999)
			result := findValueInList(3000, sortedList)
			g.Assert(result).Equal(-1)
		})

		g.It("Look for 40 in list from 1 to 102", func() {
			sortedList := makeRangedList(1, 102)
			result := findValueInList(40, sortedList)
			g.Assert(result).Equal(39)
		})

		g.It("Look for 40 in list from 1 to 102 with missing value 40", func() {
			sortedList := removeElementFromList(makeRangedList(1, 102), 39)
			result := findValueInList(40, sortedList)
			g.Assert(result).Equal(-1)
		})

		g.It("Look for 34 in list from 1 to 140", func() {
			sortedList := makeRangedList(1, 140)
			result := findValueInList(34, sortedList)
			g.Assert(result).Equal(33)
		})

		g.It("Look for 21 in non-ranged list", func() {
			sortedList := []int{19, 21, 23, 26, 34, 456, 654, 780, 800, 900}
			result := findValueInList(21, sortedList)
			g.Assert(result).Equal(1)
		})

		g.It("Look for 21 in small non-ranged list", func() {
			sortedList := []int{19, 21, 23}
			result := findValueInList(21, sortedList)
			g.Assert(result).Equal(1)
		})

		g.It("Look for 21 in non-ranged list with only 2 values", func() {
			sortedList := []int{19, 21}
			result := findValueInList(21, sortedList)
			g.Assert(result).Equal(1)
		})
	})
}
