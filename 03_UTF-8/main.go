package main

import "fmt"

func main() {
	for i := 500; i < 1000; i++ {
		fmt.Printf("%d \t %b \t %x \t %q           \n", i, i, i, i)
	}
}
