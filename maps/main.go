package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	(&jim).updateName("Jimmy")
	p := &jim
	fmt.Println(p)
	fmt.Println(&p)
	// jim.print()
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
