// Key Value Pairs

package main

import "fmt"

func main() {
     fruits := map[string]string{
	 "a":"Apple",
	 "b":"Banana",
	 "c":"Cherry",
    }	 
	
	for key, value := range fruits {
	    fmt.Printf("\n",key,value)
	} 
	}
