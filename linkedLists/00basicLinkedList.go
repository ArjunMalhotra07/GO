package main

import "fmt"

func MainFunction00() {
	myList := linkedList{}
	myList.insertBeginning(21)
	myList.insertBeginning(15)
	myList.insertBeginning(456)
	myList.insertBeginning(1895)
	myList.printList()
	myList.deleteByValue(15)
	myList.printList()
	myList.insertAfter(100, 456)
	myList.printList()
	myList.insertLast(555)
	myList.insertLast(550)
	myList.printList()
	myList.deleteByValue(550)
	myList.printList()
	myList.insertBefore(200, 456)
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

func (l *linkedList) insertBeginning(value int) {
	newNode := node{data: value}

	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
}
func (l *linkedList) insertBefore(value, before int) {
	newNode := node{data: value}

	previous := l.head

	for previous.next.data != before {

		previous = previous.next
	}
	newNode.next = previous.next
	previous.next = &newNode

}
func (l *linkedList) insertAfter(value, after int) {
	newNode := node{data: value}
	insertAfter := l.head

	for insertAfter.data != after {

		insertAfter = insertAfter.next
	}
	newNode.next = insertAfter.next
	insertAfter.next = &newNode

}

func (l *linkedList) insertLast(value int) {
	newNode := node{data: value}

	if l.head == nil {
		l.head = &newNode
		l.length++
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = &newNode
		newNode.next = nil
	}
}
func (l *linkedList) deleteByValue(value int) {
	//! If HEAD is Nil 
	if l.head == nil { //or l.length==0
		return
	}
	//! If Head is the Target Node 
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	//! If Target Node is anywhere in Between 
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

func (l *linkedList) printList() {
	if l.length == 0 {
		return
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%v ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}
