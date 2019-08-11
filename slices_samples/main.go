package main

import "fmt"

func main() {
	var s = [] int {1,2,3,4,5,6,7,8}
	fmt.Println(s)
	s2 := s[:3]
	fmt.Println(s2)
	fmt.Println("len=%d cap=%d ",len(s), cap(s))
}