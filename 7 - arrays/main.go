package main

import "fmt"

func main() {
	var age uint16 = 26
	var name string = "amin"
	var ages [3]int
	var grades [3]float64 = [3]float64{15.366, 6.3216}
	var names [2]string = [2]string{"C++", "Golang"}

	ages[0] = 10
	ages[1] = 20
	ages[2] = 30

	fmt.Println("Ages are ", ages)
	fmt.Printf("The language is %s\n", names[1])
	fmt.Println(names)
	fmt.Printf("Grades are %#v\n", grades)
	fmt.Println(ages)

	fmt.Println(name, age)

	arr := [3]int{15, 20}
	fmt.Println(arr)

	arr1 := [...]string{"C++", "golang", "java", "python"}
	fmt.Printf("Langs are %#v\n", arr1)

	fmt.Println("Size of arr is ", len(arr))
	fmt.Println(len(arr1))
}
