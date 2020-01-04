package main

import (
	"fmt"
	"io/ioutil"
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
	filename := p.Title + ".txt"
	// 함수에 command + 좌클릭 하면 함수 모양 나옴. 0600 은 rwx 리눅스 권한임.
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// 에러처리에 대한 부분은 아래와 같이 한다. 대체로 error를 같이 리턴해버리는 듯.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, error := ioutil.ReadFile(filename)
	if error != nil {
		return nil, error
	}
	// &의 경우에는 주소 참조 연산자인데, 이 주소가 어디서 나온건질 모르겠다. 리턴에 Page가 포함된다면 이것도 내부함수 취급인가..?
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("hello World")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
