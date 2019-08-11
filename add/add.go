package main

import "fmt"
func main(){
	var fname,lname string
	fmt.Println("Enter fname:")
	fmt.Scanln(&fname)
	fmt.Println("Enter lname:")
	fmt.Scanln(&lname)
	fmt.Println(fname+" ",lname)
}