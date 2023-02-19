package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	//contactInfo  // no need to declare the type def if the typeName & type are equal
}

func main() {

	// p := person{"x", "y"}
	// p := person{firstName: "x", lastName: "y"}

	// var alex person
	// fmt.Printf("%+v", alex) // %+v prints the key values in the struct person

	p := person{}
	p.firstName = "x"
	p.lastName = "y"
	fmt.Println(p)

	jim := person{
		firstName: "jim",
		lastName:  "p",
		contact: contactInfo{
			email:   "a@sd",
			zipCode: 12,
		},
	}

	// Pass by value which doesn't update
	// jim.updateName("jimmy")
	// jim.print()

	// Type 1 Pass by reference which updates
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	// jimPointer.print()

	// Type 2 Strange Case
	jim.updateName("jimmy")
	jim.print()

	// Type 3 Strange Case -- gotcha for slices updating
	mySlice := []string{"hi", "how", "are", "you"}
	updateMySlice(mySlice)
	fmt.Println(mySlice)

}

func (pointerToPerson *person) updateName(name string) {
	(*pointerToPerson).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func updateMySlice(s []string) {
	s[0] = "Bye"
}
