package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)
func fileHandling() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File Reading Error", err)
		return
	}
}
func main() {
	
	go fileHandling()
lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		if strings.Contains(line, "FileHandling") {
			lines[i] = "Use GoLang to handle files"
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("test.txt", []byte(output), 0644)
	if err != nil {
		fmt.Println("File Writing error", err)
	}
}
