package linkedlist

import "log"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

//create linked list, add element at the front
//print linked list

func (ll *LinkedList) AddElementAtFront(data int) {
	var firstNode Node
	firstNode.data = data

	if ll.head == nil {
		ll.head = &firstNode
		return
	}

	firstNode.next = ll.head
	ll.head = &firstNode
}

func (ll *LinkedList) Print() {

	if ll.head == nil {
		return
	}

	tempHead := ll.head
	for ; tempHead != nil ; {
		log.Println(tempHead.data)
		tempHead = tempHead.next
	}
}

// add at back
// delete an element given index
// delete linked list
// length
// reverse
// add element at given index
// print recursive
// search element
// detect loop
// merge two sorted linked list
