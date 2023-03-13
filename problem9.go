/* Q9) Convert this map in json*/

package main

import (
	"encoding/json"
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

	// map to json

	jsonDataResult, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonDataResult))
}
