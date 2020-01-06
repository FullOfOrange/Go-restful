# Go

https://github.com/eddycjy/go-gin-example 이곳의 프로젝트를 참고해서 예약 플랫폼 API를 뽑아보는 것을 목표로 둔다.

### BUILD
```go
go build .
```

### RUN
```bash
./go-restful
```

### Documents

공부하면서 정리한 md들을 모아놓음. 대체로 Go 만의 특징적인 것들이 적혀있을 예정

1. [Array & Slice](https://github.com/FullOfOrange/Go/blob/master/docs/arrays%26slices.md)
2. [Function literals](https://github.com/FullOfOrange/Go/blob/master/docs/functionLiterals.md)

### Reference

공부하다가 찾아냈던 유용한 사이트들을 모아놓음.

1. [GOPATH 등의 환경변수 설명](https://steemit.com/golang/@dakeshi/go-gopath-gobin)
   일단 모든 GO Project 디렉터리들은 GOPATH 하위에 있어야 한다. // echo \$GOPATH

   > 1.13 에서 변경되었다. 3. 의 내용을 참고하자.

2. [Go 언어 한국어 위키](https://github.com/golang-kr/golang-doc/wiki)
   번역된 것이 별로 없는것이 함정이다.

3. [Go Modules](https://johngrib.github.io/wiki/golang-mod/)
   1.11 부터 지원하는 공식적인 Go 모듈 관리 도구이며 Go에 종속되어있다. 이것은 GOPATH와 상관없이 동작한다.

   1.13 부터는 기몬으로 된다는데 모르겠다. 현재 이 프로젝트는 Go modules를 사용한다.
   ```go
   go mod init example.com/users/projects
   ```
   위의 명령어를 실행하면 go.mod 가 생긴다. 이 파일 내에서 각각의 의존성이 관리된다.
   만약 여기에 // indirect 라는 주석이 있으면, 다른 모듈의 디펜던시임을 의미한다.

   go install을 하면 모든 go 파일을 훑으면서 의존성을 설치하고, mod 에 추가한다. (없으면 추가하고, 있으면 있는놈 버전으로 설치함.) 설치는 GOPATH의 pkg 에 설치를 해서 관리를 한다. 나의 현재 GOPATH는 $HOME/go 이다.

   go.sum 은 각각 모듈들의 특정 버전 해쉬값이다. 이걸 비교해서 모듈이 예상치 않게 변경되는 일을 방지한다.

   이 외에는 모두 위의 링크를 따르자.

4. [Makefile](https://sohlich.github.io/post/go_makefile/)
   go 는 makefile을 이용해서 build나 test 코드를 쉽게 실행할 수 있다. 마치 package.json 의 script 같은 느낌..?

### Tips

여기있는 설명들은 내가 익숙해지면 필요가 없어지겠지만 그래도 써놓긴 함.

1. var i = 1 은 i := 1 로 대체될 수 있다.

2. \_ 의 경우에는 변수를 날려버리는 역할을 한다. (그냥 사용 안해버리는 것
   ```go
   body, _ := ioutil.ReadFile(filename)
   위의 경우에는 리턴이 (body, error) 인데, 여기서 error 를 날려버린 것.
   ```
