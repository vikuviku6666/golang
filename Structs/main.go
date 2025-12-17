package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
		firstName string
		lastName string
		contact contactInfo
	}

func main() {
	vinay := person{
		firstName: "Vin",
		lastName: "Kun",
		contact: contactInfo{
			email: "vin@gmail.com",
			zipCode: 500053,
		},
	}

	vinayPointer := &vinay
	vinayPointer.updateName("Vai")
	vinay.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
		(*pointerToPerson).firstName = newFirstName
}

// struct as receiver, it is as copy not reference to the original struct
func (pointerToPerson *person) print() {
	fmt.Printf("%+v\n", *pointerToPerson)
}

// Value Types: int, float, string, bool, structs (use pointers to change these things in a function)
// Reference Types: slices, maps, channels, pointers, functions
