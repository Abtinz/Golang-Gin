package main

import "fmt"

func main(){

	number := 1
	pointer := &number

	fmt.Println(number , pointer)

	fmt.Println(*pointer) //read from pointer

	//change  value of pointing address
	*pointer = 10
	fmt.Println(number)


} 