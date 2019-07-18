package main

import "net/http"
import "fmt"

func main () {
	links := [] string {
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",

	}
	c := make(chan string)
	for _, link := range links {
	go checkLink(link,c)
	}
	for i:=0; i<len(links); i++ {
	//for {
			//fmt.Println(<-c)
			go checkLink(<-c,c)
	}
}
func checkLink(link string, c chan string) {
_,err := http.Get(link)
if(err !=nil){
	fmt.Println(link, "might be down")
	c <- link
	return 
}
fmt.Println(link, "is up")
c <- link
}