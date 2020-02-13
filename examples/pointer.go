package main

import "fmt"

type Test struct {
	name string
	age  int
}

func test(temp []Test){
	temp[0].name = "b";
}

func main() {
	var temp []Test

	temp = append(temp,Test{"a",10})
	fmt.Println(temp[0])

	test(temp);

	fmt.Println(temp[0])
}
