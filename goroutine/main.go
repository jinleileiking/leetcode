package main

import (
	"fmt"
)

func main() {

	a := []string{"a", "b", "c", "d", "e"}
	b := []string{"1", "2", "3", "4", "5"}

	aChan := make(chan bool)
	bChan := make(chan bool)

	go func() {
		for _, v := range a {
			fmt.Println(v)
			aChan <- true
			<-bChan
		}
	}()

	go func() {
		for _, v := range b {
			<-aChan
			fmt.Println(v)

			bChan <- true
		}
	}()

	for {
	}
	return
}
