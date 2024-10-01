package main

import "fmt"

func square(num int) int {
    return num * num
}

func main() {
     var number int
	 fmt.Println("Enter a number:")
     fmt.Scan(&number)

     fmt.Print("Square", square(number))

}
