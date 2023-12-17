package main

import "fmt"

func main(){

	email := make(map[string] string)
	email["Abtin"] = "abtinzandi@gmail.com"
	email["Sarvin"] = "sarvinnami@gmail.com"
	email["Hamid"] = "hamidrezazandi@gmail.com"

	fmt.Println(email)
	fmt.Println(len(email))
	fmt.Println(email["Abtin"])

}