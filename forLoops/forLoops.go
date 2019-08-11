package main

import "fmt"
func main(){
	k:=1
	for k<=5{
	fmt.Println(k)
	k++
}
fmt.Println(add(20,30))
for range "Niha"{
	fmt.Println("Niharika")
}
strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}
	person :=map[string]string{"cn":"niha","ln":"dopp","id":"YT836"}
for index, element :=range person{
	fmt.Println("Index :", index, " Element :", element)
}
}
func add(x int, y int) (total int){
	total=x+y
	return total
}