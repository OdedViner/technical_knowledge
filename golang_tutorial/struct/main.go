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
			zipCode: 1234,
		},
	}
	jim.updateName("jimmy")
	jim.print()

	mySlice := []string{"hi", "oded"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
	mySlice = append(mySlice, "212")
	fmt.Println(mySlice)
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	fmt.Println(namePointer)
	printPointer(namePointer)

}
func updateSlice(s []string) {
	s[0] = "bye"
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
