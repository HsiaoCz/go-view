package main

import "fmt"

// go的泛型

type Slice[T int | string | float64] []T

func main() {
	var s Slice[string]
	s = append(s, "hello")
	fmt.Printf("%T\n", s)
	fmt.Println(s)
}
