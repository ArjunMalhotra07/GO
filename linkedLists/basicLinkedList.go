package main

import "fmt"

func main() {
	myList := linkedList{}
	myList.insertBefore(21)
	myList.insertBefore(15)
	myList.insertBefore(456)
	myList.insertBefore(1895)
	myList.printList()
	myList.deleteByValue(15)
	myList.printList()
	myList.insertAfter(100, 456)
	myList.printList()

}

type node struct {
	next *node
	data int
}
type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) insertBefore(value int) {
	newNode := node{data: value}

	if l.head != nil { //   45>98>14
		newNode.next = l.head //21>45>98>14
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
}

func (l *linkedList) printList() {
	if l.length == 0 { //or
		return
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%v ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}
func (l *linkedList) deleteByValue(value int) {
	if l.head == nil { //or l.length==0
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head

	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--

}
func (l *linkedList) insertAfter(value, after int) {
	newNode := node{data: value}
	insertAfter := l.head

	for insertAfter != nil {
		if insertAfter.data == after {
			newNode.next = insertAfter.next
			insertAfter.next = &newNode

		}
		insertAfter = insertAfter.next
	}

}
