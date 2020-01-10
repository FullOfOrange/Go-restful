# Panic, Defer, Recover

예외 처리에 사용될 수 있는 세가지 키워드에 대해 말한다.
이것들은 try - catch - finally 구문의 키워드들 처럼 동작하지만, 블록과 형태의 제약이 없이 자유롭다. (각각의 키워드에 위치 제약이 없음)

### Defer

defer 키워드는 특정 문장, 함수를 defer가 실행된 함숙가 return 하기 직전에 실행하도록 한다. 일반적으로는 finally 처럼 마지막 cleanup을 위해 사용하거나, 무엇인가 기다리기 위해 사용된다. finally 와 비슷하게 파일을 닫는 함수를 짜본다.

```go
package main

import "os"

func main() {
    f, err := os.Open("1.txt")
    if err != nil {
        panic(err)
    }

    // main 마지막에 파일 close 실행
    defer f.Close()

    // 파일 읽기
    bytes := make([]byte, 1024)
    f.Read(bytes)
    println(len(bytes))
}
```

### Panic

Go 에서 일반적이지 않은, 무언가 잘못된 상황이 발생하면 panic 상태라고 한다. 대체로 준비가 되어있지 않은 오류에 대해 빠르게 실패시키기 위해 (함수를 빠르게 종료시키거나 프로그램을 아예 죽일 때) 사용한다.

예를 들면 배열을 5개만 할당해놓고, 5번 인덱스를 접근하거나 하면 panic을 일으킨다.

강제로 panic 을 일으키는 방법은 다음과 같다.

```go
package main

import "os"

func main() {
    openFile("Invalid.txt")
    println("Done") //이 문장은 실행 안됨
}

func openFile(fn string) {
    f, err := os.Open(fn)
    if err != nil {
        panic(err)
    }
    // 파일 close 실행됨
    defer f.Close()
}
```
이런식으로 panic() 을 사용해서 panic을 일으킬 수 있다.

### Recover

recover() 은 패닉을 해결해주는 함수이다.
자세한 정보는 [이곳](https://cjwoov.tistory.com/8)에 있다. 햇갈리면 확인해보자.