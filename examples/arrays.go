package main

import (
	"fmt"
)
func array(){
	var a [4]int
	a[0] = 1
	variable := a[0]
	fmt.Print(variable)

	// 하단의 두개는 동일한 타입을 가진다. (컴파일러가 개수를 세어줌)
	b := [2]string{"penn", "Teller"}
	bb := [...]string{"penn", "Teller"}
	fmt.Print(b, bb)
}
func slice(){
	letters:= []string{"a","b","c","d"}
	var s []byte
	s = make([]byte,5,5)
	// s == []byte{0,0,0,0,0}
}
func main(){
	array()
	slice()
}