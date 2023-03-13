/*Enhance above function to return a map where key is input number and value is
power of 2 of the input number  */

package main

import (
	"fmt"
	"math"
)

func powerToMap(nums []int) map[int]int {
	result := make(map[int]int)
	for _, num := range nums {
		result[num] = int(math.Pow(2, float64(num)))
	}
	return result
}

func main() {
	nums := []int{10, 43, 0, 8}
	results := powerToMap(nums)
	for k, v := range results {
		fmt.Printf("%d -> %d\n", k, v)
	}
}
