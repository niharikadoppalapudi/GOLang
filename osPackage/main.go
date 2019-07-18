package main

import (
	"fmt"
	"os"
)

func main() {
	os.Chdir("C:/Users/ND074W")
	mydir, err := os.Getwd()
	if err == nil {
		fmt.Println(mydir)
	}
	fmt.Println(os.ExpandEnv("$USER lives in ${HOME}."))
	fmt.Println(os.Getegid())
	fmt.Println(os.Getpid())
	fmt.Println(os.Hostname())
}
