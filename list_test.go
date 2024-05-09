package list

import (
	"log"
	"testing"
)

func check(err error, t *testing.T) {
	if err == nil {
		return
	}
	t.Error(err)
}

func TestList(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	print := func(p Person) {
		log.Printf("%v", p)
	}

	addressBook := New[Person](OptThreadSafety)
	check(addressBook.AddLast(Person{Name: "obiwan", Age: 44}), t)
	check(addressBook.AddFirst(Person{Name: "ahsoka", Age: 33}), t)
	check(addressBook.AddLast(Person{Name: "yoda", Age: 22}), t)
	check(addressBook.AddLast(Person{Name: "boba", Age: 52}), t)
	check(addressBook.AddFirst(Person{Name: "anakin", Age: 10}), t)
	check(addressBook.Add(1, Person{Name: "padme", Age: 20}), t)
	addressBook.For(print)
	log.Printf("======================")
	ret := addressBook.GetData(func(p Person) bool {
		return p.Age > 20
	})
	for _, v := range ret {
		log.Println("get data:", v)
	}
	foundValue, err := addressBook.GetDataByIndex(1)
	if err != nil {
		t.Error(err)
	} else {
		log.Println("get data index:", foundValue)
	}
	log.Printf("empty=%v, len=%d\n", addressBook.IsEmpty(), addressBook.Len())
	check(addressBook.RemoveIndex(1), t)
	addressBook.For(print)
	log.Printf("======================")
	addressBook.RemoveFirst()
	addressBook.For(print)
	log.Printf("======================")
	addressBook.RemoveLast()
	addressBook.For(print)
	log.Printf("======================")
	addressBook.RemoveIndex(1)
	addressBook.For(print)
	log.Printf("======================")

	addressBook.RemoveAll()
	addressBook.For(print)
	log.Printf("======================")
}
