package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
	defer fmt.Println("Defer in main F'n")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 1 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))

	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	defer fmt.Println("Defer statement will be executed?")
	g(i + 1)

}
