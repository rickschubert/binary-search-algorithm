package binary_search

import (
	"math"
)

func getMiddleIdx(list []int) int {
	listLength := len(list)
	roundedMiddle := math.Round(float64(listLength) / 2)
	midIdx := int(roundedMiddle) - 1
	return midIdx
}

func isValueInFirstHalf(idx int, value int, list []int) bool {
	return list[idx] > value
}

func getListHalfWithValue(idx int, value int, list []int) ([]int, int) {
	midIdx := getMiddleIdx(list)
	isInFirstHalf := isValueInFirstHalf(midIdx, value, list)
	lengthLeftHalf := midIdx
	// cuts list either in left-hand half or right-hand half
	if isInFirstHalf {
		list = append(list[:midIdx+1], list[len(list):]...)
	} else {
		list = append(list[:0], list[midIdx:]...)
	}
	return list, lengthLeftHalf
}

func valueIsOutOfSortedListRange(value int, list []int) bool {
	if value < list[0] || value > list[len(list)-1] {
		return true
	} else {
		return false
	}
}

func makeCopy(list []int) []int {
	return append([]int(nil), list...)
}

func findValueInList(value int, list []int) int {
	// Automatically fail an empty list
	if len(list) == 0 {
		return -1
	}
	// Fail if the value is smaller or larger than minimum and maximum list values
	if valueIsOutOfSortedListRange(value, list) {
		return -1
	}
	// If above check passed and the list has only 1 item, the index must be 0
	if len(list) == 1 {
		return 0
	}

	var index int
	var firstValTempList = list[0]
	var tempList []int
	var lengthLeftHalf int
	// copy slice in order to not muate it
	tempList = makeCopy(list)

	// Infinite loop that breaks if the index was found
	for true {
		midIdx := getMiddleIdx(tempList)
		tempList, lengthLeftHalf = getListHalfWithValue(midIdx, value, tempList)
		if tempList[0] != firstValTempList {
			index += lengthLeftHalf
			firstValTempList = tempList[0]
		}
		if len(tempList) == 2 {
			if tempList[1] == value {
				index++
			} else if tempList[0] != value {
				index = -1
			}
			break
		}
	}
	return index
}
