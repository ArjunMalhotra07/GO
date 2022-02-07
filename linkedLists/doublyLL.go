package main

import (
	"fmt"
)

func main() {

	myList := linkedList_Doubly{}
	myList.insertBeginning_Doubly(21)
	myList.insertBeginning_Doubly(15)
	myList.insertBeginning_Doubly(1100)
	myList.insertBeginning_Doubly(456)
	myList.insertBeginning_Doubly(1895)
	myList.print_DoublyLL()
	myList.deleteByValue(1100)
	myList.print_DoublyLL()
	myList.insertLast(1500)
	myList.print_DoublyLL()

}

type node_Doubly struct {
	next *node_Doubly
	prev *node_Doubly
	data int
}
type linkedList_Doubly struct {
	head   *node_Doubly
	tail   *node_Doubly
	length int
}

func (l *linkedList_Doubly) insertBeginning_Doubly(value int) {
	newNode := node_Doubly{data: value}

	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
		l.length++
	} else {
		newNode.next = l.head
		l.head.prev = &newNode
		l.head = &newNode
		l.length++
	}
}

func (l *linkedList_Doubly) insertLast(value int) {
	newNode := node_Doubly{data: value}

	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
		l.length++
	} else {
		currentNode := l.tail
		currentNode.next = &newNode
		newNode.prev = currentNode
		l.tail = &newNode
	}
}
func (l *linkedList_Doubly) deleteByValue(value int) {
	if l.head == nil {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.head.prev = nil
		l.length--
		return
	}
	if l.tail.data == value {
		l.tail = l.tail.prev
		l.tail.next = nil
		l.length--
		return
	}

	currentNode := l.head
	for currentNode.data != value {
		if currentNode.next == nil {
			return
		}
		currentNode = currentNode.next
	}
	currentNode.prev.next = currentNode.next
	currentNode.next.prev = currentNode.prev
	l.length--

}

func (l *linkedList_Doubly) print_DoublyLL() {
	f := fmt.Println
	//f(l.tail.data)
	if l.length == 0 {
		return
	}
	// f()
	// f("Printing Doubly LL From Back")
	// currentNode_Back := l.tail
	// for currentNode_Back != nil {
	// 	fmt.Printf("%v ", currentNode_Back.data)
	// 	currentNode_Back = currentNode_Back.prev
	// }
	// f()
	fmt.Println("Printing Doubly LL From Front")

	currentNode_Front := l.head
	for currentNode_Front != nil {
		fmt.Printf("%v ", currentNode_Front.data)
		currentNode_Front = currentNode_Front.next
	}
	f()
}
