package main

import "fmt"

func main() {


	var greetings []string = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	slice1 := greetings[0:2]
	slice2 := greetings[1:4]
	slice3 := greetings [3:5]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	var message string = "Hi 😀 and 😁"
	var messageRune []rune = []rune(message)
	fmt.Printf("The fourth rune is: %c\n",messageRune[3])


	wilson := Employee{}
	brown := Employee{"Brown", "Brown", 6789}
	var bob = Employee {firstName: "Bob", lastName: "Smith", id: 12345}

	fmt.Println(wilson)
	fmt.Println(brown)
	fmt.Println(bob)
	
}

type Employee struct {
	firstName string
	lastName string
	id int
}