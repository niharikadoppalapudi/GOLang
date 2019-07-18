package main 

import "fmt"

func main() {
	var student map[string] string
	fruits := make(map[string]string)
	colors:= map[string]string {
		"red" : "#ff0000",
		"green" : "#4bf745",
		"orange" : "#6df42a",
	}
	colors["white"] = "#ffffff"
	delete(colors, "green")
	fmt.Println(fruits)
	fmt.Println(student)
	//fmt.Println(colors)
	printMap(colors)
}
func printMap(c map[string]string){
	for color, hex := range c {
		fmt.Println("Hex code for", color,"is", hex)
	}

}