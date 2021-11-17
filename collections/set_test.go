package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetAdd(t *testing.T) {
	var sut JCSet
	sut.Add("")
	sut.Add(0)
	sut.Add("abcde")
	want := []Item{"", 0, "abcde"}
	has := sut.Items()
	fmt.Print(want)
	fmt.Print(has)

	if !reflect.DeepEqual(has, want) {
		t.Error("Unexpectes values in set after Add() operation")
	}
}

func TestSetInterSec(t *testing.T) {
	var sut1 JCSet
	var sut2 JCSet
	var want JCSet

	sut1.Add(1)
	sut1.Add(2)
	sut1.Add("abc")

	sut2.Add(1)
	sut2.Add("abc")
	sut2.Add("abc")
	sut2.Add("interia.pl")

	want.Add(1)
	want.Add("abc")

	has := sut1.Intersection(&sut2)

	if !has.Equals(want) {
		t.Error("Unexpectes values in set intesection", "has:", has, "want: ", want)
	}
}
