package main          

import (
    "fmt"             
    "os"    
	"strings"          
)

func main() {         
	data , err  := os.ReadFiles("test.md")
	if err != nil {
		fmt.Println("error reading file : ", err)
		return
	}


	lines := strings.Split(string(data), "\n")
	for _, line : range lines {
		fmt.Println(line)
	} 

}