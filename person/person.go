package person

import "fmt"

// 3.6 Structs
type Person struct {
	name string
	age  int
}

// create receiver function
func (p Person) SayHello() {
	fmt.Printf("Hello! my name is %s and I'm %d. ", p.name, p.age)
}

func (p Person) GetKoreanAge() {
	fmt.Printf("My Korean age is %d", p.age-1)
}

// create receiver pointer function
// without *, p Person will be copied (=> unmodified).
// with *(pointer), modify the value of struct
func (p *Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age
	fmt.Println("SetDetails func: ", p)

}
