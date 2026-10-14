package dsa

type Dnode struct {
	value int
	next  *Dnode
	prev  *Dnode
}

type DoublyLinkedList struct {
	head *Dnode
	tail *Dnode
}

func (ddl *DoublyLinkedList) InsertFront(value int) {
	newNode := &Dnode{value: value}
	if ddl.head == nil {
		ddl.head = newNode
		ddl.tail = newNode
		return
	}

	newNode.next = ddl.head
	ddl.head.prev = newNode
	ddl.head = newNode
}

func (ddl *DoublyLinkedList) InsertAtEnd(value int) {
	newNode := &Dnode{value: value}
	if ddl.tail == nil {
		ddl.tail = newNode
		ddl.head = newNode
		return
	}

	newNode.prev = ddl.tail
	ddl.tail.next = newNode
	ddl.tail = newNode
}

func (ddl *DoublyLinkedList) Delete(value int) {
	if ddl.head == nil {
		return
	}

	temp := ddl.head
	for temp != nil {
		if temp.value == value {
			// only one elem
			if temp.prev == nil && temp.next == nil {
				ddl.head = nil
				ddl.tail = nil
				return
			}

			// Deleting the head
			if temp.prev == nil {
				ddl.head = temp.next
				temp.next.prev = nil
				return
			}

			// Deleting the tail
			if temp.next == nil {
				ddl.tail = temp.prev
				temp.prev.next = nil
				return
			}

			// Deleting a middle elem
			temp.prev.next = temp.next
			temp.next.prev = temp.prev
			return
		}
		temp = temp.next
	}
}
