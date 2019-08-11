package main

import (
	"testing"
//"fmt"
)
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		//fmt.Println("Expected deck length of 20, but got", len(d))
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
