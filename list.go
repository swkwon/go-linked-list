package list

import (
	"errors"
	"sync"
)

var ErrNotFoundNode = errors.New("not found node")

type Option int

const (
	OptThreadSafety Option = 1
)

type Node[T any] struct {
	prev *Node[T]
	next *Node[T]
	data T
}

type List[T any] struct {
	cnt   int
	head  *Node[T]
	tail  *Node[T]
	optTS *optThreadSafety
}

type optThreadSafety struct {
	mutex sync.Mutex
}

func New[T any](option ...Option) *List[T] {
	var optTS *optThreadSafety
	for _, v := range option {
		switch v {
		case OptThreadSafety:
			optTS = &optThreadSafety{}
		}
	}
	return &List[T]{
		optTS: optTS,
	}
}

func (l *List[T]) AddFirst(v T) error {
	l.lockImpl()
	defer l.unLockImpl()
	return l.addImpl(0, v)
}

func (l *List[T]) AddLast(v T) error {
	l.lockImpl()
	defer l.unLockImpl()
	return l.addImpl(l.cnt, v)
}

func (l *List[T]) Add(index int, v T) error {
	l.lockImpl()
	defer l.unLockImpl()
	return l.addImpl(index, v)
}

func (l *List[T]) IsEmpty() bool {
	l.lockImpl()
	defer l.unLockImpl()
	return l.cnt == 0
}

func (l *List[T]) Len() int {
	l.lockImpl()
	defer l.unLockImpl()
	return l.cnt
}

func (l *List[T]) RemoveAll() {
	pos := l.head
	for pos != nil {
		next := pos.next
		pos.prev = nil
		pos.next = nil
		pos = next
	}
	l.head = nil
	l.tail = nil
}

func (l *List[T]) RemoveFirst() error {
	l.lockImpl()
	defer l.unLockImpl()
	foundNode, err := l.findNodeImpl(0)
	if err != nil {
		return err
	}
	return l.removeNodeImpl(foundNode)
}

func (l *List[T]) RemoveLast() error {
	l.lockImpl()
	defer l.unLockImpl()
	foundNode, err := l.findNodeImpl(l.cnt - 1)
	if err != nil {
		return err
	}
	return l.removeNodeImpl(foundNode)
}

func (l *List[T]) RemoveIndex(index int) error {
	l.lockImpl()
	defer l.unLockImpl()
	foundNode, err := l.findNodeImpl(index)
	if err != nil {
		return err
	}
	return l.removeNodeImpl(foundNode)
}

func (l *List[T]) GetDataByIndex(index int) (T, error) {
	l.lockImpl()
	defer l.unLockImpl()
	ret, err := l.findNodeImpl(index)
	if err != nil {
		var retValue T
		return retValue, err
	} else {
		return ret.data, nil
	}
}

func (l *List[T]) GetData(f func(T) bool) []T {
	l.lockImpl()
	defer l.unLockImpl()
	var ret []T
	temp := l.head
	for temp != nil {
		if f(temp.data) {
			ret = append(ret, temp.data)
		}
		temp = temp.next
	}
	return ret
}

func (l *List[T]) RemoveNode(node *Node[T]) error {
	l.lockImpl()
	defer l.unLockImpl()
	return l.removeNodeImpl(node)
}

func (l *List[T]) For(f func(T)) {
	l.lockImpl()
	defer l.unLockImpl()
	temp := l.head
	for temp != nil {
		f(temp.data)
		temp = temp.next
	}
}

func (l *List[T]) lockImpl() {
	if l.optTS == nil {
		return
	}
	l.optTS.mutex.Lock()
}

func (l *List[T]) unLockImpl() {
	if l.optTS == nil {
		return
	}
	l.optTS.mutex.Unlock()
}

func (l *List[T]) removeNodeImpl(node *Node[T]) error {
	if l.cnt <= 0 {
		return ErrNotFoundNode
	}
	prev := node.prev
	next := node.next
	if l.cnt == 1 {
		if l.head == node && l.tail == node {
			l.head = nil
			l.tail = nil
		}
	} else if prev == nil { // if node is head
		l.head = next
		next.prev = nil
	} else if next == nil { // if node is tail
		l.tail = prev
		prev.next = nil
	} else {
		prev.next = next
		next.prev = prev
	}

	node.prev = nil
	node.next = nil
	l.cnt--
	return nil
}

func (l *List[T]) addImpl(index int, v T) error {
	addNode := &Node[T]{}
	addNode.data = v

	if l.cnt == 0 {
		l.head = addNode
		l.tail = addNode
	} else if index == 0 {
		l.head.prev = addNode
		addNode.next = l.head
		l.head = addNode
	} else if index == l.cnt {
		l.tail.next = addNode
		addNode.prev = l.tail
		l.tail = addNode
	} else {
		foundNode, err := l.findNodeImpl(index)
		if err != nil {
			return err
		}
		prevNode := foundNode.prev
		prevNode.next = addNode
		addNode.prev = prevNode
		addNode.next = foundNode
		foundNode.prev = addNode
	}

	l.cnt++
	return nil
}

func (l *List[T]) findNodeImpl(index int) (*Node[T], error) {
	if index < 0 || l.cnt <= index {
		return nil, ErrNotFoundNode
	}
	var ret *Node[T]
	var temp int
	if index < l.cnt/2 {
		temp = 0
		ret = l.head
		for temp != index {
			ret = ret.next
			temp++
		}
	} else {
		ret = l.tail
		temp = l.cnt - 1
		for temp != index {
			ret = ret.prev
			temp--
		}
	}
	return ret, nil
}
