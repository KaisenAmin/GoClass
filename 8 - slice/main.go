package main

import "fmt"

func main() {
	arr1 := [5]int{15, 32, 65, 95, 87}
	my_slice := []string{"Hello Go programmers", "C++", "Java"}

	fmt.Println("MY slice is ", my_slice)
	fmt.Println("Len of slice is ", len(my_slice))
	fmt.Println("cap of slice is ", cap(my_slice))
	fmt.Println(arr1)

	slice2 := make([]string, 2, 5)
	slice2[0] = "amin"
	slice2[1] = "golang"

	fmt.Println(slice2)
	fmt.Println("len of slice2 ", len(slice2))
	fmt.Println("Cap of slice2 ", cap(slice2))

	slice2 = append(slice2, "Php", "hello", "b")
	fmt.Println(slice2)

	fmt.Println("len of slice2 ", len(slice2))
	fmt.Println("Cap of slice2 ", cap(slice2))

	slice3 := arr1[1:2] // n - 1
	slice3 = arr1[1:4]  // slicing

	slice4 := arr1[:3]
	slice5 := arr1[1:]

	fmt.Println("slice5 ", slice5)
	fmt.Println("slice4 ", slice4)
	fmt.Println("slice3 ", slice3)

}
