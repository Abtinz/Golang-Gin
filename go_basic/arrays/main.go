package main

import "fmt"

func main(){
	var array [2]string
	array[0] = "Abtin"
	array[1] = "Golang"
	fmt.Println(array)

	//declare and assigning 
	new_array := [2]string{"Abtin","Gin"}
	new_array[0] = "Django"
	fmt.Println(new_array)

}