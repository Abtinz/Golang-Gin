package main

import "fmt"

func main(){

	var name = "abtin"
	var age = 21
	fmt.Println(name, age)
	fmt.Printf("%T\n",age)
	fmt.Printf("%T\n",name)

	const zandi = true

	//Shorthand
	fullname := "Abtin Zandi"
	fmt.Println(fullname)

	//specifying float types
	var first_float_type float32 = 1.3
	var second_float_type float64 = 1.3
	fmt.Printf("%T\n",first_float_type)
	fmt.Printf("%T\n",second_float_type)


}