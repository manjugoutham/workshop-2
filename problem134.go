/* Q1) Read the LearnerData.json file and print the data
https://drive.google.com/file/d/1sR4_5gHLoMyLLG0JQewO7EtDvrieuI3s/view?usp=share_link */

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var learnerData []struct {
	Name      string `json:"name"`
	EmailId   string `json:"EmailId"`
	Education struct {
		Qualification string `json:"qualification"`
		Stream        string `json:"stream"`
		YearOfPassOut string `json:"yearOfPassOut"`
	} `json:"Education"`
	LocationPreferences []string `json:"locationPreferences"`
}

func main() {

	file, err := os.Open("LearnerData.json")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&learnerData)
	if err != nil {
		fmt.Println(err)
	}
	err = json.NewDecoder(file).Decode(&learnerData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Learners:")
	// for _, learnerjson := range learnerData {
	// 	fmt.Printf("  Name is : %s\n", learnerjson.Name)
	// 	fmt.Printf("  EmailId is : %s\n", learnerjson.EmailId)
	// 	fmt.Printf("  Qualification is : %s\n", learnerjson.Education.Qualification)
	// 	fmt.Printf("  Stream is : %s\n", learnerjson.Education.Stream)
	// 	fmt.Println(" YearOfPassOut is : %s", learnerjson.Education.YearOfPassOut)
	// 	fmt.Printf("  LocationPreferences is : %s\n", learnerjson.LocationPreferences)
	// 	fmt.Println()
	// }

	/*
		Question 4 :
		From the above data, print the names and year of passout for persons who
	    passed out 2020 or later. */

	fmt.Println("Print the Names and Year of passout for persons who passed out 2020 or later:")
	for _, passedNameandYear := range learnerData {

		if passedNameandYear.Education.YearOfPassOut >= "2020" {
			fmt.Printf("  Name: %s\n", passedNameandYear.Name)
			fmt.Printf("  Year of passout: %s\n", passedNameandYear.Education.YearOfPassOut)
		}
	}

	/*  Question 5 :
	Using the above data, please print out the data (name , location preferences )
	  who have location preferences as West.  */

	fmt.Println("Print out the data (name , location preferences)who have location preferences as West :")
	for _, nameAndLocations := range learnerData {
		for _, nameandlocation := range nameAndLocations.LocationPreferences {
			if nameandlocation == "West" {
				fmt.Printf("  Name : %s\n", nameAndLocations.Name)
				fmt.Printf("  Location Preferences : %v\n", nameAndLocations.LocationPreferences)
			}
		}
	}
}
