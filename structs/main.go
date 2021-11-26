package main

import "fmt"

type Person struct {
	firstName string
	lastName string
	age int
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func main()  {
	p := Person{firstName: "Nagesh", lastName: "Dhope", age: 29}

	fmt.Println(p.FullName())
}
