package main

import (
	"fmt"

	"github.com/chowanij/algorithms-in-golang/collections"
)

func main() {
	var mySet collections.JCSet
	fmt.Println("Algorithms and data structures in golang")
	mySet.Add("abcd")
	mySet.Add(123)
	mySet.Add(123.123)
	fmt.Println(mySet.Pop())
	fmt.Println(mySet.Pop())
	fmt.Println(mySet.Pop())
	fmt.Println(mySet.Pop())
}
