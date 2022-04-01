package main

import "fmt"

func main() {
	fmt.Println("teste")
	fmt.Println(Soma(10, 10))
}

func Soma(a int, b int) int {
	return a + b
}
