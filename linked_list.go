package dsa

type Lnode struct {
	value int
	next  *Lnode
}

type LinkedList struct {
	head *Lnode
}

func (l *LinkedList) InsertAtBegining(value int) {
	newNode := &Lnode{value: value}
	if l.head == nil {
		l.head = newNode
	}

	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) InsertAtLast(value int) {
	newNode := &Lnode{value: value}
	if l.head == nil {
		l.head = newNode
	}

	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode
}

func (l *LinkedList) Delete(value int) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
	}

	temp := l.head
	for temp.next != nil && temp.next.value != value {
		temp = temp.next
	}

	if temp.next != nil {
		temp.next = temp.next.next
	}
}
