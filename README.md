# Go

golang study repo

### Reference

1. [GOPATH 등의 환경변수 설명](https://steemit.com/golang/@dakeshi/go-gopath-gobin)
   - 일단 모든 GO Project 디렉터리들은 GOPATH 하위에 있어야 한다. // echo \$GOPATH

### Directory Structure

### Tips

여기있는 설명들은 내가 익숙해지면 필요가 없어지겠지만 그래도 써놓긴 함.

1. var i = 1 은 i := 1 로 대체될 수 있다.

2. \_ 의 경우에는 변수를 날려버리는 역할을 한다. (그냥 사용 안해버리는 것
   ```go
   body, _ := ioutil.ReadFile(filename)
   위의 경우에는 리턴이 (body, error) 인데, 여기서 error 를 날려버린 것.
   ```
