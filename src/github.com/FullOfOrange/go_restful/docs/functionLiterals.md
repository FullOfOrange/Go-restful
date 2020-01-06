# Function literals

익명함수를 표현하는 법이 Function literals 임.

```go
FunctionLit = "func" Signature FunctionBody .
```

사용법은 아래처럼 이름이 없는 function 을 정의하면 된다.
```go
func(a, b int, z float 64) bool { return a*b < int(z) }
```

Function literals는  변수에 대입되거나, 즉시 실행될 수 있음.
즉시 실행 함수는 표현법이 js 처럼 마지막에 () 로써 실행을 표현해줘야함.
```go
f := func(x, y int) int { return x + y }
func(ch chan int) { ch <- ACK }(replyChan)
```

Function literals 는 Closures 이다. 주변 함수에 의해, function literal이 접근될 수 있는 한, (변수로써 살아있을 경우) 이것은 계속해서 접근할 수 있게 된다.