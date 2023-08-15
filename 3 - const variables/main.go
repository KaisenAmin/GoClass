package main

import "fmt"

const Pi float32 = 3.1565
const Data = 1561

// var data int

// just like enum in other language like c or c++
const (
	Red = iota // ioata + 1
	White
	Black
)

const (
	FirstName string = "amin"
	age              = 27
)

const (
	baseprice  = 100
	discount   = 10
	finalPrice = baseprice * (baseprice - discount) / 100
)

var dataString string = "in global "

func main() {
	const FirstName string = "golang"
	// var dataString string = "in local"

	fmt.Printf("the final price is %v and the type is %T", finalPrice, finalPrice)
	fmt.Println(dataString)
	fmt.Printf("Pi is %v\n", Pi)
	fmt.Printf("Pi type is %T\n", Pi)

	fmt.Println(FirstName)
	fmt.Println(Data)
	// Pi = 1565.2156
	fmt.Println(Pi)

	fmt.Println(Red, White, Black)

}
