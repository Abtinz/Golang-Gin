package main

import "fmt"

func sumation()  func (int) int  {
	sum := 0 
	return func (x int) int  {
		sum += x
		return sum
	}
}

func main(){

	i := 0
	sum := sumation()
	for i < 5{
		fmt.Println(sum(i))
		i ++
	}

} 