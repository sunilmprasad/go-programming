package main

import "fmt"

func main() {
	fmt.Println("here iam cumming")

	foo()

	fmt.Println("something foo")

	for i := 0; i < 50; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {

	fmt.Println("am foo too")
}

func bar() {
	fmt.Println("pothum poda")
}

//control flow
//(1)sequence
//(2)loop;iterative
//(3)conditional
