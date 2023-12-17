package main

import "fmt"

func greeting(name string) string{
	return "hi " + name
}

func sum(num1 , num2 int){
	fmt.Println(num1 + num2)
}

func main(){
	fmt.Println(greeting( "Abtin"))
	sum(2 , 2 )
}