package main

import ( 
	"fmt"
	"strings"
	"os"
	)

func WordstoNum(s string) string {

	var num string
	temp:= strings.Split(s," ")
	for i:=0;i<len(temp);i++ {
		if strings.EqualFold(temp[i],"zero") {
			num = num + "0"
		} else if strings.EqualFold(temp[i],"one") {
			num = num + "1"
		} else if strings.EqualFold(temp[i],"two") {
			num = num + "2"
		} else if strings.EqualFold(temp[i],"three") {
			num = num + "3"
		} else if strings.EqualFold(temp[i],"four") {
			num = num + "4"
		} else if strings.EqualFold(temp[i],"five") {
			num = num + "5"
		} else if strings.EqualFold(temp[i],"six") {
			num = num + "6"
		} else if strings.EqualFold(temp[i],"seven") {
			num = num + "7"
		} else if strings.EqualFold(temp[i],"eight") {
			num = num + "8"
		} else if strings.EqualFold(temp[i],"nine") {
			num = num + "9"
		} else {
			return "Invalid Input"
			}
	}
	return num
}

func main() {
	fmt.Println(WordstoNum(os.Args[1]))
}
		
		
