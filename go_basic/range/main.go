package main

import "fmt"

func main(){
	ids := []int{1,2,33,21,45,66,745,3}
	for i , id:=range ids{
		fmt.Println(i,id)
	}
} 