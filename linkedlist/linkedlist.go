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

func (ll *LinkedList) AddElementAtEnd(data int) {
	node := Node{data: data, next: nil}

	if ll.head == nil {
		ll.head = &node
		return
	}

	tempHead := ll.head
	for ; tempHead.next != nil; {
		tempHead = tempHead.next
	}
	tempHead.next = &node
}

func (ll *LinkedList) Print() {

	if ll.head == nil {
		return
	}

	tempHead := ll.head
	for ; tempHead != nil; {
		log.Println(tempHead.data)
		tempHead = tempHead.next
	}
}

func (ll *LinkedList) DeleteAtGivenIndex(index int) {
	if ll.head == nil {
		return
	}

	if index == 1 {
		ll.head = ll.head.next
		return
	}
	tempHead := ll.head
	for i := 1; i < index-1 && tempHead.next != nil; i++ {
		tempHead = tempHead.next
	}
	if tempHead.next == nil {
		return
	}
	tempHead.next = tempHead.next.next
}

func (ll *LinkedList) Delete() {
	ll.head = nil
}

func (ll *LinkedList) Length() int {
	length := 0
	tempHead := ll.head
	for ; tempHead != nil; {
		tempHead = tempHead.next
		length++
	}
	return length
}

func (ll *LinkedList) Reverse() {
	if ll.head == nil || ll.head.next == nil {
		return
	}

	var prevNode *Node
	currentNode := ll.head
	nextNode := ll.head

	for ; nextNode.next != nil; {
		nextNode = currentNode.next
		currentNode.next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}
	currentNode.next = prevNode
	ll.head = currentNode
}

func (ll *LinkedList) AddElementAtGivenIndex(data int, index int) {
	if ll.head == nil {
		return
	}

	newNode := Node{data: data, next: nil}
	tempHead := ll.head

	if index == 1 {
		newNode.next = tempHead
		ll.head = &newNode
		return
	}
	for i := 1; i < index-1 && tempHead != nil; i++ {
		tempHead = tempHead.next
	}

	if tempHead == nil {
		return
	}
	newNode.next = tempHead.next
	tempHead.next = &newNode
}

//done
//add at front
//traverse
// add at back
// delete an element given index
// delete linked list
// length
// reverse
// add element at given index


//todo
// print recursive
// search element
// detect loop
// merge two sorted linked list
