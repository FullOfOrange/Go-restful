package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// Global 에서 쓰이는 Must들은 Runtime이나 CompileTime에 조건에 맞지 않으면 panic을 일으키는 것들임.
// template 들을 캐싱하는 것이 좋음. 어차피 쓰일거 계속해서 io를 하지말고 메모리에 올려두자.
var templates = template.Must(template.ParseFiles("../templates/edit.html", "../templates/view.html"))

// validation 을 저장하는 path를 정의.
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte
}

// function literal & closure
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invaild Page Title")
	}
	return m[2], nil
}

// Reciever는 특정 structure 의 메소드라는 뜻임.
// 여기에서 (p *Page) 는 포인터 리시버인데,
// 값을 전달받을 때 하드카피를 할지 포인터를 받을지를 결정함.
func (p *Page) save() error {
	// 이 함수는 데이터를 저장하는 용도로 사용됨.
	filename := "../data/" + p.Title + ".txt"
	// 함수에 command + 좌클릭 하면 함수 모양 나옴. 0600 은 rwx 리눅스 권한임.
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// 에러처리에 대한 부분은 아래와 같이 한다. 대체로 error를 같이 리턴해버리는 듯.
func loadPage(title string) (*Page, error) {
	filename := "../data/" + title + ".txt"
	body, error := ioutil.ReadFile(filename)
	if error != nil {
		return nil, error
	}
	// &의 경우에는 주소 참조 연산자인데, 이 주소가 어디서 나온건질 모르겠다. 리턴에 Page가 포함된다면 이것도 내부함수 취급인가..?
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, temp string, p *Page) {
	err := templates.ExecuteTemplate(w, temp+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// form value가 존재하지 않을 경우도 생각해줘야할 것 같음.
	body := r.FormValue("body")
	// 형 변환은 []byte() 이런식으로 하는 것 같음.
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}
