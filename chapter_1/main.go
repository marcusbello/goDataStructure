package main

import "fmt"

func main() {
	fmt.Println("Hello GoWorld Data structures and Algorithms")

	//Binary Search Exercise
	list := []int{1, 2, 3, 4, 5, 6}
	v := 4
	binarySearch := binarySearch(list, v)
	fmt.Println(binarySearch)

	//Sequential Search Exercise
	list = []int{3, 2, 1, 4, 6, 5}
	v = 4
	y := 7
	sequentialSearch1 := sequentialSearch(list, v)
	sequentialSearch2 := sequentialSearch(list, y)
	fmt.Println(sequentialSearch1)
	fmt.Println(sequentialSearch2)

	//Sum of Array Exercise
	list = []int{7, 2, 3, 4, 5}
	sumArray := sumArray(list)
	fmt.Println(sumArray)
}

func binarySearch(list []int, v int) bool {
	size := len(list)
	var middle int
	start := 0
	end := size - 1
	for start <= end {
		middle = start + (end-start)/2
		if list[middle] == v {
			return true
		} else {
			if list[middle] < v {
				start = middle + 1
			} else {
				end = middle - 1
			}
		}
	}
	return false
}

func sequentialSearch(list []int, v int) bool {
	for _, value := range list {
		if value == v {
			return true
		}
	}
	return false
}

func sumArray(list []int) int {
	/*
		Write a method that will return the sum of all the elements of the integer list,
		given list as an input argument.
	*/
	size := len(list)
	sum := 0
	for i := 0; i < size; i++ {
		sum = sum + list[i]
	}
	return sum
}
