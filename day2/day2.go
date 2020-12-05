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
		var min, max int64
		var character, password string
    	if i >= 0 {
			// var err2 error
			min, _ = strconv.ParseInt(each_ln[0:i], 10, 64)
			fmt.Println(min)
        	j := strings.Index(each_ln, " ")
        	if j >= 0 {
				max, _ = strconv.ParseInt(each_ln[i+1 : j], 10, 64)
         	   	fmt.Println(max)
			}
			z := strings.Index(each_ln, ": ")
			character = each_ln[j+1:z]
			fmt.Println(character)
			password = each_ln[z+2:len(each_ln)]
			fmt.Println(password)
			res := strings.Count(password, character)
			fmt.Println(res)
			if int64(res) >= min && int64(res) <= max {
				valid++
			}
		}
		
	} 
	fmt.Println("Number of valid passwords:", valid)
	
}

