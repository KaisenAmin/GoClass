package main

import "fmt"

func main() {
	// bitwise not

	a := 10
	b := 3

	// var name string = "Hello amin" // assignment operator

	// a -> 1010
	// b -> 0011
	// &^ > 1000
	// 8

	bitwiste := a &^ b
	fmt.Printf("%d &^ %d = %d\n", a, b, bitwiste)

	// relationals operator
	// > < >= <= == !=

	var one, two int
	fmt.Printf("Enter number one : ")
	fmt.Scan(&one)

	fmt.Printf("Enter number two : ")
	fmt.Scan(&two)

	grater := one > two
	fmt.Printf("%d > %d == %t\n", one, two, grater)

	less := one < two
	fmt.Printf("%d < %d == %t\n", one, two, less)

	graterEqual := one >= two
	fmt.Printf("%d >= %d == %t\n", one, two, graterEqual)

	lessEqual := one <= two
	fmt.Printf("%d <= %d == %t\n", one, two, lessEqual)

	equality := one == two
	fmt.Printf("%d == %d == %t\n", one, two, equality)

	notEqual := one != two
	fmt.Printf("%d != %d == %t\n", one, two, notEqual)

	var number1 float32 = 10.0
	var number2 float32 = 10.0

	fmt.Println(number1 == number2)
	fmt.Println("amin" > "hello")
	fmt.Println("amin" == "amin")
}
