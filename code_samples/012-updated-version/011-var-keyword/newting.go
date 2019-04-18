package main

import ("fmt")

var y = 04
func main(){
	x := 15
	fmt.Println(x)
	fmt.Println(y)
	foo()
}
func foo(){
	fmt.Println("again:" ,y)
}
