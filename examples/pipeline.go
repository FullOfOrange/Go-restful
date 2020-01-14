package main

import (
	"fmt"
	// "math/rand"
	"time"
)

// func factoryRand(count int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		for i := 0; i < count; i++ {
// 			out <- rand.Intn(100)
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func printChanNum(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func(){
// 		for i := range in {
// 			fmt.Println(i)
// 			out <- i
// 		}
// 		close(out);
// 	}()
// 	return out
// }
// func main() {
// 	// 작업시작시간 기록
// 	start := time.Now()

// 	sum := 0
// 	for n := range printChanNum(factoryRand(1000000) {
// 		sum += n
// 	}
// 	fmt.Printf("Total Sum of Squares: %d\n", sum)

// 	// 작업 종료 후 시간기록
// 	elapsed := time.Since(start)
// 	fmt.Println("작업소요시간: ", elapsed)
// }

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			time.Sleep(5 * time.Second)
			c1 <- "one"
		}
	}()

	go func() {
		for {
			time.Sleep(10 * time.Second)
			c2 <- "two"
		}
	}()

	for {
		fmt.Println("start select------------------")
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		default: fmt.Println("default")
		}
		fmt.Println("end select-------------------\n\n")
	}
}
