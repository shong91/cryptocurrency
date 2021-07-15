package main

import "fmt"

// seldomly use const
const name1 string = "const value"

func plus(a ...int) int {
	// can multiple return
	// return a + b, name

	// can iterate 
	var total int
	for _, item := range a {
		total += item 
	}

	return total
}


func main() {
	// init variable
	var name string = "hong"
	// shortcut is only availiable inside of function 
	age := 13
	fmt.Println("it works! ", name, age)

	// int / uint(unsinged integer) 
	// unit cannt be negative number. 

	result := plus(2,2,3,4,5,6,7)
	fmt.Println(result)

	name = "hong ! ! ! ! 1 ! ! ! ! my name is"
	for _, letter := range name {
		fmt.Println(string(letter))
	}

}