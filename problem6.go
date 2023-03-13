/* Write a function which computes Power of 2 for a number. Function takes an integer as an argument and returns power of 2 for that integer.
Develop this function without using the math package of Golang.
Write a main function which takes user input and computes the power of 2 using the above function and prints the same on terminal in the following way :
“Power of 9 is 512”
Test this function with at least few integer numbers like 9, 0, 7, 1, 2   */

package main

import (
	"fmt"
)

func powerOfTwo(number int) int {
	result := 1
	for i := 0; i < number; i++ {
		result = result * 2
	}
	return result
}

func main() {
	var number int
	fmt.Print("Enter the number : ")
	fmt.Scan(&number)
	powervalue := powerOfTwo(number)
	fmt.Printf("Power of %d value is %d\n", number, powervalue)
}
