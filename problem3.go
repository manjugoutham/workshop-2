/*From the above data create a map of name and year of passout.*/

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
	fmt.Println(learnerData)

	// for _, learnerjson1 := range learnerData {
	// fmt.Printf("  Name is : %s\n", learnerjson.Name)
	// fmt.Printf("  EmailId is : %s\n", learnerjson.EmailId)
	// fmt.Printf("  Qualification is : %s\n", learnerjson.Education.Qualification)
	// fmt.Printf("  Stream is : %s\n", learnerjson.Education.Stream)
	// fmt.Println(" YearOfPassOut is : %s", learnerjson.Education.YearOfPassOut)
	// fmt.Printf("  LocationPreferences is : %s\n", learnerjson.LocationPreferences)
	// fmt.Println(learnerjson1)
	// }

	mapjson1 := make(map[string]string)
	for _, learnerData1 := range learnerData {
		mapjson1[learnerData1.Name] = learnerData1.Education.YearOfPassOut
	}
	// fmt.Println(mapjson)

	/* From the above data, print the names and year of passout for persons who
	passed out 2020 or later.  */

}
