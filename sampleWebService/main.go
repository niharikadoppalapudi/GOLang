package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Result struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
	Type    string `json:"type"`
	Value   Value  `json:"Value"`
}

type Value struct {
	ID         int      `json:"id"`
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}

func getRandomName(url string, ch1 chan<- io.ReadCloser, swg *sync.WaitGroup) {
	defer swg.Done()
	name, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
	}
	ch1 <- name.Body

}
func getRandomJoke(url string, ch1 chan<- io.ReadCloser, swg *sync.WaitGroup) {
	defer swg.Done()
	joke, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	ch1 <- joke.Body
}
func main() {
	handler := func(res http.ResponseWriter, req *http.Request) {
		var t Result
		ch1 := make(chan io.ReadCloser)
		var swg sync.WaitGroup
		swg.Add(2)
		go getRandomName("http://uinames.com/api/", ch1, &swg)
		go getRandomJoke("http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]", ch1, &swg)
		_ = json.NewDecoder(<-ch1).Decode(&t)
		_ = json.NewDecoder(<-ch1).Decode(&t)
		swg.Wait()
		json.NewEncoder(res).Encode(t)
	}
	http.HandleFunc("/combine", handler)
	http.ListenAndServe(":5000", nil)

}
