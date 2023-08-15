package main

import "fmt"

// global scope
// package scope

// FirstName can use in other packages
var firstName string = "golang programmers"
var age rune = -326

// lastName := "Hello" // infer string dont do this in global scope

var (
	grade float64
	data  = 236
)

func main() {
	// this is local scope
	fmt.Println(firstName)
	var name string = "amin" // local variable

	// fmt.Println(lastName)

	fmt.Println("the age is", age)
	fmt.Printf("The name is %s\n", name)

	fmt.Printf("First name before call functions is [%s]\n", firstName)
	changeGlobalVariable() // function call
	fmt.Printf("First name After call function is [%s]\n", firstName)

	fmt.Printf("%v, %v", grade, data)

}

func changeGlobalVariable() {
	firstName = "This is my second name"
}
