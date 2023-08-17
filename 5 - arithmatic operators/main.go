package main

import "fmt"

func main() {
	var a, b int

	fmt.Printf("Enter a and b please : ")
	fmt.Scan(&a, &b)

	addition := a + b + 10 // a + b ->
	subtraction := a - b*2
	division := float32(a) / float32(b)
	multiplication := a * b / 2
	remainder := a % b
	bitwiseNot := a &^ b // ampersand not

	// 10 -> 1010
	// 3 ->  0011
	// 1010
	// &^
	// 0011
	// 1000

	fmt.Printf("addition of %d + %d = %d\n", a, b, addition)
	fmt.Printf("subtraction of %d - %d = %d\n", a, b, subtraction)
	fmt.Printf("divition of %d / %d = %f\n", a, b, division)
	fmt.Printf("multiplication of %d * %d = %d\n", a, b, multiplication)
	fmt.Printf("remainder of %d %% %d = %d\n", a, b, remainder)
	fmt.Printf("bitwisenot is %d &^ %d = %d\n", a, b, bitwiseNot)

}
