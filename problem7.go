/* Enhance the Power of 2 function using the math package of GoLang. Note: Power function of Math package :
func Pow(x, y float64) float6é4 Write a function which takes a series (slice) of numbers as an input argument
and returns a series (slice) of numbers which are Power of 2 numbers of respective input numbers. This function uses the Power of 2 function developed
Write a main function where a slice of integers are defined. Compute the power of this slice using the above function.
For example:
For Input: [10 43 0 8]
Output will be :
[1024 16 8 1 256]    */

package main

import (
	"fmt"
	"math"
)

func powerToSlice(nums []int) []int {
	result := make([]int, len(nums))
	for i, number := range nums {
		result[i] = int(math.Pow(2, float64(number)))
	}
	return result
}

func main() {
	nums := []int{10, 43, 0, 8}
	results := powerToSlice(nums)
	fmt.Println(results)
}
