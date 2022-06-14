package main

import "fmt"

var list []int
var v int
var k int
var y int

func main() {
	fmt.Println("Hello GoWorld!\nData structures and Algorithms In GO\n")

	//Factorial
	v = 4
	fmt.Println("4 factorial is:", factorial(v))

	//Max sum of subarray which is the maximum
	list = []int{1, -1, -3, 2, -4, 5, -13, 8}
	fmt.Println("Max sub array sum:", maxSubArraySum(list))

	//Rotate elements in an array K number of times
	list = []int{10, 20, 30, 40, 50, 60}
	k = 3
	fmt.Println("Rotate elements in K times:", rotateArray(list, k))

	//Binary Search Exercise
	list = []int{1, 2, 3, 4, 5, 6}
	v = 4
	binarySearch := binarySearch(list, v)
	fmt.Println("Binary Search Array:", binarySearch)

	//Sequential Search Exercise
	list = []int{3, 2, 1, 4, 6, 5}
	v = 4
	y = 7
	sequentialSearch1 := sequentialSearch(list, v)
	sequentialSearch2 := sequentialSearch(list, y)
	fmt.Println("Sequential Search for value v in array:", sequentialSearch1)
	fmt.Println("Sequential Search for value y in array:", sequentialSearch2)

	//Sum of Array Exercise
	list = []int{7, 2, 3, 4, 5}
	sumArray := sumArray(list)
	fmt.Println("Sum of all items in an array:", sumArray)
}

func factorial(v int) int {
	if v <= 1 {
		return 1
	}
	return v * factorial(v-1)
}

func maxSubArraySum(list []int) int {
	n := len(list)
	maxSoFar := 0
	maxEnding := 0
	for i := 0; i < n; i++ {
		maxEnding = maxEnding + list[i]
		if maxEnding < 0 {
			maxEnding = 0
		}
		if maxSoFar < maxEnding {
			maxSoFar = maxEnding
		}
	}
	return maxSoFar
}

func rotateArray(list []int, k int) []int {
	n := len(list)
	reverseArray(list, 0, k-1)
	reverseArray(list, k, n-1)
	reverseArray(list, 0, n-1)
	return list
}

func reverseArray(list []int, start int, end int) {
	i := start
	j := end
	for i < j {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
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
