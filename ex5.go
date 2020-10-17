package main

import (
	"fmt"
	"os"
	"strconv"
	)

func main() {
	
	str1 := os.Args[1]
	operation := os.Args[2]
	str2 := os.Args[3]
	
	var bitsize int
	complex1,_ := strconv.ParseComplex(str1,bitsize)
	complex2,_ := strconv.ParseComplex(str2,bitsize)
	
	if operation=="+" {
		fmt.Println(complex1+complex2)
	} 
	if operation=="-" {
		fmt.Println(complex1-complex2)
	}
	if operation=="*" {
		fmt.Println(complex1*complex2)
	}
}
	