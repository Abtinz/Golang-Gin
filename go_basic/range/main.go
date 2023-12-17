package main

import "fmt"

func main(){
	ids := []int{1,2,33,21,45,66,745,3}

	//loop over array
	for index , id:=range ids{
		fmt.Println(index,id)
	}

	//no need to returned data just like python
	for _ , id:=range ids{
		fmt.Printf("id: %d\n",id)
	}

	//RANGE OVER MAP
	code := map[string] int{"Baeca" : 1 , "Man united" : 2}
	for key , value:=range code{
		fmt.Println( key , value)
	}

} 