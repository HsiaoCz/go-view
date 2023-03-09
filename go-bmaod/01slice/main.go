package main

import "fmt"

func main() {

	s := []int{1, 2, 3}
	f(s)
	fmt.Println(s)
}

func f(s []int) []int {
	s = append(s, 10)
	return s
}
