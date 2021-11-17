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

	var sut1 collections.JCSet
	var sut2 collections.JCSet

	sut1.Add(1)
	sut1.Add("abc")

	sut2.Add(1)
	sut2.Add("abc")

	booool := sut1.Equals(sut2)

	fmt.Println("booool: ", booool)

	intersection := sut1.Intersection(&sut2)

	fmt.Println("Intersection: ", intersection)


}
