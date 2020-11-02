package review

import "fmt"

func MidFind(a []int, target int) bool {
	for true {
		l := len(a)
		if l == 1 {
			return a[0] == target
		}
		limit := l / 2
		if a[limit] == target {
			fmt.Println(limit)
			return true
		} else if a[limit] > target {
			a = a[:limit]
		} else {
			a = a[limit+1:]
		}
	}
	return false
}

func binarySearch(sortedArray []int, lookingFor int) int {
	var low int = 0
	var high int = len(sortedArray) - 1
	for low <= high {
		var mid int = low + (high-low)/2
		var midValue int = sortedArray[mid]
		if midValue == lookingFor {
			return mid
		} else if midValue > lookingFor {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

//0, 1, 2, 3, 5, 6, 7
func FindClose(a []int, target int) int {
	if len(a) == 0 {
		return -1
	}
	if a[0] > target {
		return -1
	}
	var low = 0
	var high = len(a) - 1
	var result int
	for low <= high {
		mid := low + (high-low)/2
		if a[mid] >= target {
			high = mid - 1
			result = high
		} else {
			low = mid + 1
			result = mid
		}
	}
	return result
}
