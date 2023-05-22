package main

import "fmt"

func main() {
	// https://www.youtube.com/watch?v=qyM8Pi1KiiM
	// CHANNEL
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

	// Done Channel from child goroutine to parent
	// -----------------------------------------------------
	// done := make(chan bool)

	// go doWork(done)

	// time.Sleep(time.Second * 3)

	// close(done)
	// -----------------------------------------------------

	// //input
	// nums := []int{2, 3, 4, 1, 7}
	// //stage 1
	// dataChannel := sliceToChannel(nums)
	// //stage 2
	// finalChannel := sq(dataChannel)
	// //stage 3
	// for n := range finalChannel {
	// 	fmt.Println(n)
	// }

	queue := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
			fmt.Println("sent")
		}
		close(queue)
	}()

	for value := range queue {
		fmt.Println(value)
	}

	// queue := make(chan int)
	// done := make(chan bool)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		queue <- i
	// 		fmt.Println("sent")
	// 	}

	// 	done <- true
	// }()

	// for {
	// 	select {
	// 	case v := <-queue:
	// 		fmt.Println(v)
	// 	case <-done:
	// 		fmt.Println("done")
	// 		return
	// 	}
	// }

	// go func() {
	// 	ch <- 100
	// 	fmt.Println("sent")
	// }()

	// ch <- 1
	// ch <- 2

	// close(ch)

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println("Done")
}

// func doWork(done <-chan bool) {
// 	for {
// 		select {
// 		case <-done:
// 			return
// 		default:
// 			fmt.Println("...DOING")
// 		}
// 	}
// }

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
