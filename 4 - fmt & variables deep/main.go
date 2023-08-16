package main

import "fmt"

var (
	name      string  = "Hello"
	age       int     = 36
	grade     float32 = 156.151
	dataGrade float64 = 2652.15615
	flag      bool    = true // false
	status    uint    = 1651
	oneByte   int8    = -12
	n         uint8   = 255
	newOne    uint16  = 165
	newTwo    uint32  = 1651651633
	newThree  uint64  = 126651416541654
	newFout   int32   = 2165165
	ch        rune    = '\u0045'
)

type Test struct {
	x, y int
}

func main() {
	fmt.Println("Hello World", name)
	fmt.Println(n, ch, status, newOne)
	fmt.Print(15, 15641, newFout, "\n")
	fmt.Println("Hello World")

	fmt.Printf("Hello %q", name)
	fmt.Printf("Hello\tone\n")
	fmt.Printf("the grade is %f\n", dataGrade)
	fmt.Printf("the new grade is %.2f\n", dataGrade)
	fmt.Printf("new new one %8.2f\n", dataGrade)
	fmt.Printf("boolean %v\n", flag)

	fmt.Printf("|%-6d|%6d", 15, 20)
	fmt.Printf("The new string is %.2s\n", name)

	test := Test{15, 16}

	fmt.Printf("%v\n", test)
	fmt.Printf("%+v\n", test)
	fmt.Printf("%#v\n", test)
	fmt.Printf("Please tell me the type of struct %T\n", test)

	var myName string
	var myAge uint16

	fmt.Printf("Please Enter your name and age : ")
	// fmt.Scan(&myName, &myAge)
	fmt.Scanf("%s %d", &myName, &myAge)

	fmt.Printf("My name is %s and iam %d years old\n", myName, myAge)
	sentence := fmt.Sprintf("Hello my name is %s", myName)
	newSentnce := fmt.Sprintln("Hello World", 15)

	fmt.Println(newSentnce)
	fmt.Printf("The sentence is [%s]\n", sentence)

}
