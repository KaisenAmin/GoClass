package main

import "fmt"

func main() {
	var sentence = "Hello go programmers"
	var age = 26 // infer
	var myString string = "The string is so good"
	var data int = 523
	var unsignedNum uint16 = 25632
	var myByte uint8 = 62 // type byte = uint8
	var floatNumber float32 = 15.326
	var gratFloat float64 = 15.365156
	var myChar rune = 'a'

	// var name string = "Hello"
	name := "amin or Kaisen" // infer string data type
	grade := float32(15.326)

	var (
		firstName        = "amin"
		lastName  string = "tahmasebi"
		myAge     int    = 26
	)

	fmt.Printf("My name is %s and my lastName is %s and myAge is %d\n", firstName, lastName, myAge)
	fmt.Printf("The grade number is %f and the type of this grade is %T\n", grade, grade)
	fmt.Println("My name is", name)
	fmt.Printf("My emotion is like this %c\n", myChar)
	fmt.Println(floatNumber, " ", gratFloat)
	fmt.Println(myByte)
	fmt.Println("The unsigned number is", unsignedNum)
	fmt.Println(data)
	fmt.Println("The string is", myString)
	fmt.Print("My age is", age, "\n")
	// fmt.Println()
	fmt.Println(sentence)

}
