package main

import "fmt"

func main() {
	a := 401
	if a%50 == 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("No")
	}
}
