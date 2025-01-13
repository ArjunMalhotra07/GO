package linkedList

import "fmt"

type nodeStruct struct {
	data            int
	nextNodeAddress *nodeStruct
}

func MainFunction02() {
	newListVariable := addToLLFromArray([]int{5, 10, 15, 99})
	newListVariable.printLL()
	cycleLinkedL := makeCycleLinkedList([]int{10, 20, 30, 40})
	fmt.Println(checkIfCycleExist(newListVariable))
	fmt.Println(checkIfCycleExist(cycleLinkedL))

}
func (l *nodeStruct) printLL() {
	current := l
	for current != nil {
		fmt.Println(current.data)
		current = current.nextNodeAddress
	}
}

func addToLLFromArray(nums []int) *nodeStruct {
	var newList *nodeStruct = &nodeStruct{}
	for i := 0; i < len(nums); i++ {
		currentNode := newList
		newNode := &nodeStruct{data: nums[i], nextNodeAddress: nil}
		for currentNode.nextNodeAddress != nil {
			currentNode = currentNode.nextNodeAddress
		}
		currentNode.nextNodeAddress = newNode
		newNode.nextNodeAddress = nil
	}
	return newList.nextNodeAddress
}
func checkIfCycleExist(head *nodeStruct) bool {
	slow := head
	fast := head
	for fast != nil && fast.nextNodeAddress != nil {
		fast = fast.nextNodeAddress.nextNodeAddress
		slow = slow.nextNodeAddress
		if slow == fast {
			return true
		}

	}
	return false
}
func makeCycleLinkedList(nums []int) *nodeStruct {
	cycleList := &nodeStruct{}
	for i := 0; i < len(nums); i++ {
		currentNode := cycleList
		newNode := &nodeStruct{data: nums[i], nextNodeAddress: nil}
		for currentNode.nextNodeAddress != nil {
			currentNode = currentNode.nextNodeAddress
		}
		currentNode.nextNodeAddress = newNode
		newNode.nextNodeAddress = nil
	}
	cNode := cycleList
	for cNode.nextNodeAddress != nil {
		cNode = cNode.nextNodeAddress
	}
	nodeB := cycleList
	for nodeB.data != 20 {
		nodeB = nodeB.nextNodeAddress
	}
	cNode.nextNodeAddress = nodeB
	return cycleList.nextNodeAddress
}
