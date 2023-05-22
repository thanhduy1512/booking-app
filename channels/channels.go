package main

import (
	"fmt"
	"time"
)

func main() {
	// https://www.youtube.com/watch?v=qyM8Pi1KiiM
	// -----------------------------------------------------
	// myChannel := make(chan string)
	// mySecondChannel := make(chan string)

	// go func() {
	// 	myChannel <- "data"
	// }()

	// go func() {
	// 	mySecondChannel <- "cow"
	// }()

	// select {
	// case msgFromMyChannel := <-myChannel:
	// 	fmt.Println(msgFromMyChannel)
	// case msgFromAnotherChannel := <-mySecondChannel:
	// 	fmt.Println(msgFromAnotherChannel)
	// }
	// -----------------------------------------------------

	// -----------------------------------------------------
	// charChannel := make(chan string, 3)
	// chars := []string{"a", "b", "c"}

	// for _, s := range chars {
	// 	select {
	// 	case charChannel <- s:
	// 	}
	// }

	// close(charChannel)

	// for result := range charChannel {
	// 	fmt.Println(result)
	// }
	// -----------------------------------------------------

	// -----------------------------------------------------
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)

	close(done)
	// -----------------------------------------------------

}

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("...DOING")
		}
	}
}
