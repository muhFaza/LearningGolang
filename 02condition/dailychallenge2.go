package main

import "fmt"

func main() {
	fmt.Println("DAILY CHALLENGE 2")

	// NUMBER 1
	fmt.Println("NUMBER 1")
	type Person map[string]string

	var people = []Person{
		{
			"name": "Hank",
			"Age":  "50",
			"Job":  "Polisi",
		},
		{
			"name": "Heisenberg",
			"Age":  "52",
			"Job":  "Ilmuwan",
		},
		{
			"name": "Skyler",
			"Age":  "48",
			"Job":  "Akuntan",
		},
	}

	for _, v := range people {
		fmt.Printf("Hi Perkenalkan, Nama saya %v, umur saya %v, dan saya bekerja sebagai %v \n", v["name"], v["Age"], v["Job"])
	}

	// NUMBER 2
	fmt.Println("NUMBER 2")
	var rows1 int = 5
	for rows1 > 0 {
		fmt.Println("*")
		rows1--
	}

	// NUMBER 3
	fmt.Println("NUMBER 3")
	var rows3 int = 5
	for i := 0; i < rows3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
