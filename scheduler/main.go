package main

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("I am runnning task.")
}
func task2() {
	fmt.Println("I am runnning task2.")
}

func main() {

	//gocron.Every(5).Seconds().Do(task)
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(task)
	s.Every(5).Seconds().Do(task2)
	<-s.Start()
	//<-gocron.Start()
}
