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
		fmt.Println(false)
	}

	switch("Abtin"){
		case "Abtin" : 
			fmt.Println("hi abnzandi")
		case "other" :
			 fmt.Println("who are you!!!")
	}

}