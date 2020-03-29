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

func (ll *LinkedList) PrintRecursively() {
	if ll == nil {
		return
	}
	tempHead := ll
	printRecursively(tempHead.head)
}

func printRecursively(head *Node) {
	if head == nil {
		return
	}
	log.Println(head.data)
	printRecursively(head.next)
}

/**
ll having loop, eg:

0->1->2->3->4->5
      |		   |
      8<-7<-6<--

*/

func (ll *LinkedList) HasLoop() bool {
	if ll == nil || ll.head.next == nil {
		return false
	}

	fastPointer := ll.head
	slowPointer := ll.head

	for ; fastPointer != nil && slowPointer != nil; {
		if fastPointer == slowPointer && fastPointer != ll.head && slowPointer != ll.head {
			return true
		}
		fastPointer = fastPointer.next.next
		slowPointer = slowPointer.next
	}
	return false
}

func (ll *LinkedList) SearchElement(value int) *Node {
	tempHead := ll.head
	for ; tempHead != nil; {
		if tempHead.data == value {
			return tempHead
		}
		tempHead = tempHead.next
	}
	return nil
}

//done
// Create linked list, add element at the front
// Print linked list
// Add at front
// Traverse
// Add at back
// Delete an element given index
// Delete linked list
// Length
// Reverse
// Add element at given index
// Print recursive
// Detect loop
// search element

//TODO:
// merge two sorted linked list
// Nth node from the end of a Linked List
// Print the middle of a given linked list
// Write a function that counts the number of times a given int occurs in a Linked List
// Detect loop in a linked list
// Find length of loop in linked list
// Function to check if a singly linked list is palindrome
// Remove duplicates from a sorted linked list
// Remove duplicates from an unsorted linked list
// Swap nodes in a linked list without swapping data
// Pairwise swap elements of a given linked list
// Move last element to front of a given Linked List
// Intersection of two Sorted Linked Lists
// Intersection point of two Linked Lists.
// QuickSort on Singly Linked List
// Segregate even and odd nodes in a Linked List
