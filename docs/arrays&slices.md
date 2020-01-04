## Arrays / Slices

Slice 는 Array의 하위타입임. 따라서 Array가 가진 성격을 다 가지고 가면서 추가적인 뭔가가 있음

### Arrays

일반적인 언어의 배열과 동일함.

1. Array는 정의할 때 length 와 type을 정해야함.
   단, 여기서의 배열의 length가 fixed 일 경우, 이 길이는 이것의 타입중 하나임.
   ([4]int, [5]int 는 서로 다른 타입임.)

2. Array 는 알아서 모든 값들이 0으로 초기화됨

3. Go 에서는 Array는 Value로 취급된다. C에서의 Pointer 의 개념으로 보는것이 아니라 전체 배열이 변수로써 존재한다.
   따라서 이 변수를 전달할 경우, Pointer가 아니라 배열 전체가 복사된다. 이것을 주의해야함.
   이걸 피하고 싶으면, 배열의 Pointer값을 전달해줘야한다.

### Slices

Arrays는 자기의 공간을 가지고 있지만, 배열을 늘리기에는 inflexible 하다.
Slices 는 그것의 단점을 없애주고 배열보다 훨씬 강력하고 편리하다.

1. slice는 []T 라는 타입을 가진다. (T는 type) 그리고 길이를 명시할 필요가 없다.
   정의는 마치 Array 정의하듯 하면 된다.

   make로도 만들 수 있다.

   ```
   func make([]T, len, cap) []T 의 형태임.
   ```

2. slice가 정의는 되어있으나 할당이 안되어있으면 nil 이다. 이것은 length 가 0 인 slice 처럼 행동한다.

   ```go
   var i []int // === nil
   ```
   
3. slice 는 말그대로 잘라낼 수도 있다.

   ```go
   b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
   // b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
   ```

   콜론으로 인텍스를 분리해서 사용도 가능. 1 <= : < 4 이다. 메모리 공간은 동일한 곳임.

   ```go
   // b[:2] == []byte{'g', 'o'}
   // b[2:] == []byte{'l', 'a', 'n', 'g'}
   // b[:] == b
   ```

4. Slice 의 내부 공간은

   ptr (\*Elem) // slice pointer
   len (int) // length of elements
   cap (int) // maximun length of segment

   로 구성되어있음.

   ![](https://blog.golang.org/go-slices-usage-and-internals_slice-struct.png)

   기본적으로 이 slice의 변수는 가장 앞부분의 값을 가르키고, s[2:4] 같은 것을 해서 넘겨주는 것은, pointer 의 위치나 length capacity 등이 다 바뀐다. 아래의 그림을 확인하자.

   ![](https://blog.golang.org/go-slices-usage-and-internals_slice-2.png)

   따라서 아래와 같은 형태의 변경을 가지게 된다.

   ```go
   d := []byte{'r', 'o', 'a', 'd'}
   e := d[2:]
   // e == []byte{'a', 'd'}
   e[1] = 'm'
   // e == []byte{'a', 'm'}
   // d == []byte{'r', 'o', 'a', 'm'}
   ```

5. Slice 의 Capacity를 늘리는 행위는 무조건 새로운 할당을 한 다음에 이뤄져야 한다. 아래의 모든 행위들은 전부 이런 slice들의 용량을 늘리는 방법을 알려주는 것이다.

   다른 dynamic array 를 가진 언어를 흉내낸 것이다. (기본적으로 공간이 부족해지면 배열을 두배씩 뻥튀기 해준다.)

   ```go
   t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
   for i := range s {
           t[i] = s[i]
   }
   s = t // 단순히 pointer를 대입해준다. null을 안해도 되네
   ```

   이런 루핑을 해주는 copy라는 함수가 내장되어 있다.

   ```go
   func copy(dst, src []T) int
   // Copy를 사용하면 아래와 같이 make를 한 뒤에 looping을 안해도 된다.
   t := make([]byte, len(s), (cap(s)+1)*2)
   copy(t, s)
   s = t
   ```

   아래는 특정 slice 뒤에 data를 추가해주는 함수이다.

   ```go
   // 여기서 data ...byte는 byte 형태의 여러개의 형태를 받을 수 있다는 뜻인듯 이걸 배열 형태로 반환해주고
   // js에서 ...values 랑 비슷한 느낌
   func AppendByte(slice []byte, data ...byte) []byte {
       m := len(slice)
       n := m + len(data)
       if n > cap(slice) { // if necessary, reallocate
           // allocate double what's needed, for future growth.
           newSlice := make([]byte, (n+1)*2)
           copy(newSlice, slice)
           slice = newSlice
       }
       slice = slice[0:n]
       copy(slice[m:n], data)
       return slice
   }
   ```

   그리고 위의 함수도 대신 해주는 append라는 함수가 있다.

   ```go
   func append(s []T, x ...T) []T
   
   // ### 아래처럼 사용하면 됨.
   a := make([]int, 1)
   // a == []int{0}
   a = append(a, 1, 2, 3)
   // a == []int{0, 1, 2, 3}
   
   // ### 만약 배열 하나를 덧붙이고 싶으면 copy보다는 array... 을 사용해서 붙여보자. ...는 구조분해를 해준다.
   a := []string{"John", "Paul"}
   b := []string{"George", "Ringo", "Pete"}
   a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
   // a == []string{"John", "Paul", "George", "Ringo", "Pete"}
   ```

   nil 은 length 가 0 인 slice 처럼 행동하기에 다음과 같은 함수도 정의가 가능하다.

   ```go
   // Filter returns a new slice holding only
   // fn으로 들어온 함수의 실행 결과에 만족한 elements 만 새로운 slice에 할당하는 행위를 할 수 있다.
   func Filter(s []int, fn func(int) bool) []int {
       var p []int // == nil
       for _, v := range s {
           if fn(v) {
               p = append(p, v)
           }
       }
       return p
   }
   ```

6. 메모리 관리

   Slices 로 불러온 전체 항목중에 slices 된 작은 항목만 필요하지만, 잘못된 사용으로  인해 전체 Slices가 메모리에 올라가 있는 경우가 있다. (pointer 를 리턴할 때 전체 항목을 가르키는 친구를 리턴해버리면 GC가 이것을 수거 못할 때도 있다.. 는 것 같음.)

   ```go
   var digitRegexp = regexp.MustCompile("[0-9]+")
   
   func FindDigits(filename string) []byte {
       b, _ := ioutil.ReadFile(filename)
       return digitRegexp.Find(b)
   }
   ```

   위 같은 경우, return이 전체 파일이 포함된 []byte가 리턴된다. 이런 경우가 메모리 수거가 안되는 경우이다. 따라서 이를 해결하기 위해서 아래와 같이 짜면 좋다.

   ```go
   func CopyDigits(filename string) []byte {
       b, _ := ioutil.ReadFile(filename)
       b = digitRegexp.Find(b)
       c := make([]byte, len(b))
       copy(c, b)
       return c
   }
   ```

   

#### Reference

https://blog.golang.org/go-slices-usage-and-internals