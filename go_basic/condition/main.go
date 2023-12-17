package main

import "fmt"

func main(){
	num1 := 1
	num2 := 2

	if(num1 > num2){
		fmt.Println(true)
	}else if(num1 == num2){
		fmt.Println("equal")
	}else{
		fmt.Print(false)
	}

}