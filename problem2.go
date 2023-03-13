// /* Save the above json data in a suitable structure */

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type LearnerData1 struct {
	Name                string     `json:"Name"`
	EmailId             string     `json:"EmailId"`
	Education1          Education1 `json:"Education"`
	LocationPreferences []string   `json:"Location Preferences"`
}

type Education1 struct {
	Qualification string `json:"Qualification"`
	Stream        string `json:"Stream"`
	YearOfPassOut string `json:"Year Of PassOut"`
}

func main() {

	learnerDatajson := []LearnerData1{
		{
			Name:    "Rammohan",
			EmailId: "mram@blsite.com",
			Education1: Education1{
				Qualification: "BE",
				Stream:        "Computer Science",
				YearOfPassOut: "2022",
			},
			LocationPreferences: []string{"South"},
		},
	}
	file, err := os.Create("learnerData1.json")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	//write the data
	err = encoder.Encode(learnerDatajson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("done")

}
