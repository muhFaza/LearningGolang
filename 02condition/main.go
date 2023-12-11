package main

import "fmt"

func main() {
	fmt.Println("CONDITIONAL")
	
	age := 0

	if age >= 18 {
		fmt.Println("First if, age =", age)
	} else  if age < 18 && age > 0 {
		fmt.Println("Else if age < 18 =", age)
	} else {
		fmt.Println("Else age is <= 0, age =", age)
	}

	switch {
		case age >= 18:
			fmt.Println("Switch > 18, age =", age)
		case age < 18 && age > 0:
			fmt.Println("Switch 1-18, age =", age)
		default: 
			fmt.Println("Default, age =", age)
	}

	fmt.Println("LOOP 1")
	for i := age; i < age + 10; i++ {
		fmt.Printf("%v ",i)
	}
	
	fmt.Println("\nLOOP 2")
	iteration := 0
	for iteration < 10 {
		fmt.Printf("%v ", iteration)
		iteration++
	}

	// ARRAY
	fmt.Println("\n ARRAY")
	array := [3]int{1,2,3}
	fmt.Println(array)

	multidimension := [2][3]int{
		{1,2,3},
		{4,5,6},
	}
	fmt.Println(multidimension)

	// SLICE
	slice := []string{"Apple", "Manggo", "Banana"}

	slice = append([]string{"Orange"}, slice...)
	fmt.Println(slice)

	// MAP
	var obj = map[string]string{
		"name" : "Faza",
		"email" : "faza@gmail.com",
	}
	value, exist := obj["name"]
	fmt.Println(value, exist)

	fmt.Println(obj)

	// SLICE OF MAP || ARRAY OF OBJECTS
	var aob = []map[string]string{
		{"name": "gery"},
		{"name": "bagas"},
	}
	fmt.Println(aob)
}