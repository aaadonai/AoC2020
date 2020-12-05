package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
  
    if err != nil { 
        log.Fatalf("failed to open") 
  
	} 
	
	scanner := bufio.NewScanner(file) 

	scanner.Split(bufio.ScanLines) 
	var text []string 
	
	for scanner.Scan() { 
        text = append(text, scanner.Text()) 
	} 
	
	file.Close() 

	var valid int
	for _, each_ln := range text { 
		fmt.Println(len(each_ln))
		i := strings.Index(each_ln, "-")
		var first, second int64
		var character, password string
    	if i >= 0 {
			first, _ = strconv.ParseInt(each_ln[0:i], 10, 64)
			fmt.Println(first)
        	j := strings.Index(each_ln, " ")
        	if j >= 0 {
				second, _ = strconv.ParseInt(each_ln[i+1 : j], 10, 64)
         	   	fmt.Println(second)
			}
			z := strings.Index(each_ln, ": ")
			character = each_ln[j+1:z]
			fmt.Println(character)
			password = each_ln[z+2:len(each_ln)]
			fmt.Println(password)
			fmt.Println("Password first:", password[first-1:first])
			fmt.Println("Password second:", password[second-1:second])
			if password[first-1:first] == character && password[second-1:second] != character {
				valid++
				fmt.Println("true")
			} else if password[first-1:first] != character && password[second-1:second] == character {
				valid++
				fmt.Println("true")
			}
			
		}
		
	} 
	fmt.Println("Number of valid passwords:", valid)
	
}

