package main

import ( 
	"fmt"
	"os"
	"strings"
	)

func NumtoWords(s string) string {
	var words string
	temp:= strings.Split(s,"")
	for k:=0;k<len(s);k++ {
		if s[k]>='a' && s[k]<='z' {
			return "Not a number"
		}
	}

	for i:=0;i<len(temp);i++ {
		if strings.Compare(temp[i],"0")==0 {
			words = words + "zero "
		}
		if strings.Compare(temp[i],"1")==0 {
			words = words + "one "
		}
		if strings.Compare(temp[i],"2")==0 {
			words = words+ "two "
		}
		if strings.Compare(temp[i],"3")==0 {
			words = words + "three "
		}
		if strings.Compare(temp[i],"4")==0 {
			words = words + "four "
		}
		if strings.Compare(temp[i],"5")==0 {
			words = words + "five "
		}
		if strings.Compare(temp[i],"6")==0 {
			words = words + "six "
		}
		if strings.Compare(temp[i],"7")==0 {
			words = words + "seven "
		}
		if strings.Compare(temp[i],"8")==0 {
			words = words + "eight "
		}
		if strings.Compare(temp[i],"9")==0 {
			words = words + "nine "
		}
		if strings.Compare(temp[i]," ")==0 {
			continue;
		} 

				
		

}
		return words	
}
		
func main() {

for p:=1;p<len(os.Args);p++ {
		fmt.Printf(NumtoWords(os.Args[p]))
	

	}
}