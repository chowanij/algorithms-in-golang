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
