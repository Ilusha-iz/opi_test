package main

import "fmt"

func main() {
	s := []int{0, 1, 2}
	fmt.Printf("%v, %v, %v\n", s, len(s), cap(s))
	copy(s, s[1:])
	fmt.Printf("%v, %v, %v\n", s, len(s), cap(s))

}
