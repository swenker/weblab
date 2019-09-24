package test

//This test is to explore some usages of golang

import (
	"fmt"
	"testing"
)

func TestSprintf(t *testing.T) {
	renderedString := fmt.Sprintf("this is a %% and %s", "sentence")

	expected := "this is a % and sentence"
	if renderedString != expected {
		t.Error("not equal!!!!")
	}
}

func TestStringOps(t *testing.T) {
	// s := ""

	var ip *int
	okk := 10
	ip = &okk

	fmt.Printf("-------------------&ip=%d,ip=%d,*ip=%d\n", &ip, ip, *ip)
}
