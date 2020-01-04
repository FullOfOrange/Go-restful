package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

type Page struct {
	Title string
	// []byte 는 'a byte slice' 를 의미한다.
	// slice 는 Array 타입 위에서 추상화된 타입이며 뭔가가 더 있다.
	//
	// 자세한 내용은 https://blog.golang.org/go-slices-usage-and-internals
	Body []byte
}

// Reciever는 특정 structure 의 메소드라는 뜻임.
// 여기에서 (p *Page) 는 포인터 리시버인데,
// 값을 전달받을 때 하드카피를 할지 포인터를 받을지를 결정함.
func (p *Page) save() error {
	// 이 함수는 데이터를 저장하는 용도로 사용됨.
	filename := "../data" + p.Title + ".txt"
	// 함수에 command + 좌클릭 하면 함수 모양 나옴. 0600 은 rwx 리눅스 권한임.
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// 에러처리에 대한 부분은 아래와 같이 한다. 대체로 error를 같이 리턴해버리는 듯.
func loadPage(title string) (*Page, error) {
	filename := "../data" + title + ".txt"
	body, error := ioutil.ReadFile(filename)
	if error != nil {
		return nil, error
	}
	// &의 경우에는 주소 참조 연산자인데, 이 주소가 어디서 나온건질 모르겠다. 리턴에 Page가 포함된다면 이것도 내부함수 취급인가..?
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// 관련 내용은 examples/server.go 에 존재한다.
	// 아래의 len("/view/") 를 통해 /view/ 를 날려버릴 수 있다.
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<body><h1>%s</h1><p>%s</p></body>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
