package main

import (
	"fmt"
	"time"
)

func emit(chanChannel chan chan string, done chan bool) {
	//words := []string{"the", "quikc", "brown", "fox"}
	wordChannel := make(chan string)
	chanChannel <- wordChannel      // tell the channel aout the creaated channel
	defer close(wordChannel)

	words := []string{"the", "quikc", "brown", "fox"}
	i := 0

	t := time.NewTimer(2 * time.Second)

	for {

		select {
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-done:
			fmt.Printf("got Done Message\n")
			//close(done)
			done <- true // channels are bi-directional
			// bad if I took out.. program will block
			//fmt.Printf("got Done Message\n")
			//close(wordChannel)
			return

		case <-t.C:
			fmt.Printf("\n\nTimer Fired\n\n")
			return
		}

	}
	//fmt.Printf("\nCompleted\n")
	//
	//close(wordChannel)

}


func main() {

	channelCh := make(chan chan string)
	doneChannel := make(chan bool)

	go emit(channelCh, doneChannel)

	wordch := <-channelCh

	for word := range wordch {
		fmt.Printf("word: %s\t\t", word)
	}

	fmt.Printf("\n\n")

}
