package main

import (
	"fmt"

	"github.com/shong91/cryptocurrency/person"
)

// seldomly use const
const name1 string = "const value"

// 3.2 Functions
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
	// 3.1 Variables in Go
	// init variable
	var name string = "hong"
	// shortcut is only availiable inside of function
	age := 13
	fmt.Println("it works! ", name, age)

	// int / uint(unsinged integer)
	// unit cannt be negative number.

	// 3.2 Functions
	result := plus(2, 2, 3, 4, 5, 6, 7)
	fmt.Println(result)

	name = "hong ! ! ! ! 1 ! ! ! ! my name is"
	for _, letter := range name {
		fmt.Println(string(letter))
	}

	// 3.3 fmt
	x := 405940594059
	// sprintf: fmt and print
	xAsBinary := fmt.Sprintf("%b\n", x)
	fmt.Println(x, xAsBinary)

	fmt.Printf("%o\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%U\n", x)

	// 3.4 slices and arrays
	// 1. array: has certain length
	foods := [3]string{"potato", "pizza", "pasta"}
	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}
	// 2. slices: is infinite, unlimitted
	drinks := []string{"cola", "cider", "juice"}
	fmt.Printf("%v\n", drinks)
	drinks = append(drinks, "beer")
	fmt.Printf("%v\n", drinks)

	// 3.5 Pointers
	a := 2
	b := a // copy the value
	a = 12
	fmt.Println(b, a)

	c := 2
	d := &c // copy the data(memory address)
	c = 10
	fmt.Println(*d, c) // *b: print the value in memory address

	// 3.6 Structs
	// 3.7 Structs with Pointers
	kim := person.Person{}
	kim.SetDetails("kim", 33)
	fmt.Println("main func: ", kim)
	kim.SayHello()
	kim.GetKoreanAge()

}
