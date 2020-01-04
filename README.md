# Go

### Documents
공부하면서 정리한 md들을 모아놓음. 대체로 Go 만의 특징적인 것들이 적혀있을 예정

1. [Array & Slice](https://github.com/FullOfOrange/Go/blob/master/docs/arrays%26slices.md)
2. [Function literals](https://github.com/FullOfOrange/Go/blob/master/docs/functionLiterals.md)
### Tips

여기있는 설명들은 내가 익숙해지면 필요가 없어지겠지만 그래도 써놓긴 함.

1. var i = 1 은 i := 1 로 대체될 수 있다.

2. \_ 의 경우에는 변수를 날려버리는 역할을 한다. (그냥 사용 안해버리는 것
   ```go
   body, _ := ioutil.ReadFile(filename)
   위의 경우에는 리턴이 (body, error) 인데, 여기서 error 를 날려버린 것.
   ```

### Reference

공부하다가 찾아냈던 유용한 사이트들을 모아놓음.

1. [GOPATH 등의 환경변수 설명](https://steemit.com/golang/@dakeshi/go-gopath-gobin)
   일단 모든 GO Project 디렉터리들은 GOPATH 하위에 있어야 한다. // echo \$GOPATH
   
2. [Go 언어 한국어 위키](https://github.com/golang-kr/golang-doc/wiki)
   번역된 것이 별로 없는것이 함정이다.

3. [Go 기본 세팅하기](https://github.com/golang-kr/golang-doc/wiki/Go-코드를-작성하는-방법)
   Go는 이상하다. 무조건 $GOPATH 내에 모든 것이 존재해야한다. 대체로 src내에 모든 project를 관리한다고 한다...