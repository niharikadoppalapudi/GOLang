package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
}
