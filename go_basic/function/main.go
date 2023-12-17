package main

import "fmt"

func greeting(name string) string{
	return "hi " + name
}

func main(){
	fmt.Println(greeting( "Abtin"))

}