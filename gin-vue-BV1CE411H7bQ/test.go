package main

import "fmt"

func main() {

	for a := range []int{212, 123, 1231, 4, 12, 4} {
		fmt.Println(a)
	}
}
