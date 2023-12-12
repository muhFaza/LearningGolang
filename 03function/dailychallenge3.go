package main

import (
	"fmt"
	"unicode"
)

func main() {
	word := []string{"Hello", "world", "zzz"}

	fmt.Println("\nCHALLENGE 1")
	result := AlayGen(word...)
	fmt.Println(result)

	n := 7
	fmt.Println("\nCHALLENGE 2")
	fibo := Fibonacci(n)
	fmt.Println("Bilangan Fibonacci ke", n, "=", fibo)

}

func AlayGen(input ...string) string {
	var result string
	for _, val := range input {
		var tmp string
		for _, wrd := range val {
			switch unicode.ToLower(wrd) {
			case 'e':
				tmp += "3"
			case 'a':
				tmp += "4"
			case 'o':
				tmp += "0"
			case 'l':
				tmp += "1"
			case 'i':
				tmp += "!"
			default:
				tmp += string(wrd)
			}
		}
		result += tmp + " "
	}
	return result
}

func Fibonacci(n int) int {
	a, b, c := 0, 1, 0
	n++
	for i := 0; i < n; i++ {
		if i > 2 {
			// Ganjil replace a
			if i%2 == 1 {
				a = c
			}
			// Genap replace b
			if i%2 == 0 {
				b = c
			}
		}
		if i == 0 || i == 1 {
			c = i
		} else {
			c = a + b
		}
	}
	return c
}
