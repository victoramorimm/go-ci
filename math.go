package main

import "fmt"

func main() {
	fmt.Println(soma(112, 10))
	fmt.Println(soma(10, 10))
}

func soma(a int, b int) int {
	return a + b
}