# go-linked-list
go-linked-list는 goroutine으로부터 안전한 이중 연결 리스트 입니다.
# 시작하기
```
$ go get github.com/swkwon/go-linked-list@latest
```
go-linked-list는 옵션에 따라 goroutine으로 부터 안전할 수도 있고, 그렇지 않을 수 있습니다.
```
	type Person struct {
		Name string
		Age  int
	}

	addressBook := New[Person](OptThreadSafety)
```
`New[T]` 함수로 연결리스트를 초기화 할 수 있습니다. 옵션 상수인 OptThreadSafety를 입력해 주면 goroutine으로 부터 안전합니다. 필요 없을 경우 파라메타를 입력하지 않으면 됩니다.
```
    addressBook.AddLast(Person{Name: "obiwan", Age: 44})
    addressBook.AddFirst(Person{Name: "ahsoka", Age: 33})
    addressBook.Add(1, Person{Name: "padme", Age: 20})`
```
요소를 추가하기 위해서는 세가지 방법이 있습니다.
`AddLast`는 리스트의 마지막 요소가 됩니다. `AddFirst`는 첫번째 요소가 됩니다. `Add`함수는 index를 지정하여 입력할 수 있습니다.
```
	print := func(p Person) {
		log.Printf("%v", p)
	}

    addressBook.For(print)
```
`For`함수를 이용하여 함수객체를 넘겨주면 모든 요소에 접근하여 명령을 실행할 수 있습니다.
```
	ret := addressBook.GetData(func(p Person) bool {
		return p.Age > 20
	})

    foundValue, err := addressBook.GetDataByIndex(1)
```
리스트의 요소를 꺼내오는 방법은 두가지가 있습니다. `GetData`는 조건 함수객체를 넘겨주어 해당하는 모든 요소들을 slice로 꺼내옵니다. `GetDataByIndex`함수는 리스트의 인덱스에 해당하는 단일 요소만 꺼내옵니다.
```
    addressBook.RemoveIndex(1)
    addressBook.RemoveFirst()
    addressBook.RemoveLast()
    addressBook.RemoveAll()
```
삭제방법은 총 네가지가 있습니다.