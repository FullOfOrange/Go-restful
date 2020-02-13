package main

import (
	"fmt"
	"math/rand"
	"time"
)

func factoryRand(count int) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(100)
		}
		close(out)
	}()
	return out
}

func printChanNum(in <-chan int) <-chan int {
	out := make(chan int)
	go func(){
		for i := range in {
			fmt.Println(i)
			out <- i
		}
		close(out);
	}()
	return out
}
func main() {
	// 작업시작시간 기록
	start := time.Now()

	sum := 0
	for n := range printChanNum(factoryRand(1000000) {
		sum += n
	}
	fmt.Printf("Total Sum of Squares: %d\n", sum)

	// 작업 종료 후 시간기록
	elapsed := time.Since(start)
	fmt.Println("작업소요시간: ", elapsed)
}