package binarysearch

import "fmt"

// SearchInts performs a binary search for key in list and returns the index if
// key is found. Returns `-1` if key is not in list.
func SearchInts(list []int, key int) (index int) {
	return binarySearch(list, key, 0, len(list))
}

func binarySearch(list []int, key int, start int, end int) (index int) {
	if len(list) == 0 {
		return -1
	}
	if end < start {
		return -1
	}
	if start >= len(list) {
		return -1
	}
	if start == end {
		if key == list[start] {
			return start
		} else {
			return -1
		}
	}
	middle := ((end - start) / 2) + start
	fmt.Printf("binarySearch(%v, %v, %v, %v)\n", list, key, start, end)
	if key == list[middle] {
		return middle
	} else if key > list[middle] {
		return binarySearch(list, key, middle+1, end)
	} else if key < list[middle] {
		return binarySearch(list, key, start, middle)
	}
	fmt.Printf("unreachable\n")
	return -1
}
