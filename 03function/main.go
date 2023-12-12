package main

import "fmt"

func main() {
	// function
	fmt.Println(plus(3, 4))
	// function with multiple returns
	fmt.Println(plusAndMultiply(4, 4))
	plus, mult := plusAndMultiply(5,5)
	fmt.Println(plus, mult)

	//variadic function
	fmt.Println(mergeNames("faza", "bagas", "gery", "dani"))
	names := []string{"mondi", "faris", "lemon", "tomat"}
	front, back := mergeNames("bagas", names...)
	fmt.Println("MERGE FRONT:", front)
	fmt.Println("MERGE BACK:", back)


	

	var me = User{
		postCount: 1,
		name: "faza",
	}

	me.IncrementPostCount("lmao")
	
	fmt.Println(users)
}

// STRUCTS AND METHODS
type User struct {
	postCount int
	name string
}

func (u User) IncrementPostCount (msg string) string{
	return fmt.Sprintf(msg)
}

func plus(x int, y int) int {
	return x + y
}

func plusAndMultiply(x int, y int) (int, int) {
	sumOfPlus := x + y
	sumOfMultiply := x * y
	return sumOfPlus, sumOfMultiply
}

func mergeNames(person string, people ...string) ([]string, []string) {
	mergeFront := append([]string{person}, people...)
	mergeBack := append(people, person)
	return mergeFront, mergeBack
}