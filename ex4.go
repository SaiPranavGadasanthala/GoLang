package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
	)
func num2hex(num int) string {
	return strconv.FormatInt(int64(num),16)
}

func main() {
	
	var str string
	var str1 string
	var c int = 1
	if (len(os.Args)>2) {

	str1 = os.Args[1]
	str = os.Args[2]
	} else {
		str = os.Args[1]
	}
	for k:=0;k<len(str);k++ {
		if str[k]>='0' && str[k]<='9' {
			c = 1
		} else {
			c = 0
			break;
			}
		}
	if(c==1) {
		
	i, _ := strconv.Atoi(str)
	if (len(os.Args)>2 && strings.Compare(str1,"-u")==0) {
		fmt.Println(strings.ToUpper(num2hex(i)))
	} else if (len(os.Args)>2 && strings.Compare(str1,"-u")!=0) {
		
		fmt.Println("Invalid option")
		} else {
			fmt.Println(num2hex(i))
		}
} else {
	fmt.Println("Invalid number")
}
	
}
	 