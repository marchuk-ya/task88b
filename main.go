package main

import (
	"fmt"
)

//Поменять порядок цифр числа n на обратный.

func Reader() string{
	fmt.Print("Enter integer: ")
	var input string
	fmt.Scanln(&input)
	return input
}

func Reverse(n string) string {
	runes := []rune(n)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println("After reverse: ", string(runes))
	return string(runes)
}

func main() {
	n := Reader()
	Reverse(n)
}
