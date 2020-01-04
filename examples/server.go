package main

import (
	"fmt"
	"log"
	"net/http"
)

// http.HandlerFunc 타입이며 아래와 같은 인자를 가져야한다.
func handler(w http.ResponseWriter, r *http.Request) {
	// fprintf 는 다른 함수와 동일하게 첫 인자로 들어온 친구한테 out을 해주는 것.
	fmt.Fprintf(w, "Hi %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
