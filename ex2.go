package main

import (
	"fmt"
	"os"
	"math"
	"strconv"
	)

func maxPower(m,n int) int {
	var i int
	for i=0;int(math.Pow(float64(n),float64(i)))<=m;i++ {
		//Loop will iterate untill the condition is false
		}
		return i-1
	}
						
					
	
	

func main() {
	
	num1,_:= strconv.Atoi(os.Args[1])
	num2,_:= strconv.Atoi(os.Args[2])
	
	fmt.Println(maxPower(num1,num2))
}