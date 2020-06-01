package linkedlist

import "log"

/**

Definition: Circular linked list is a linked list where all nodes are connected to form a circle. There is no NULL at the end. A circular linked list can be a singly circular linked list or doubly circular linked list.

Advantages of Circular Linked Lists:
- Any node can be a starting point. We can traverse the whole list by starting from any point. We just need to stop when the first visited node is visited again.
- Useful for implementation of queue. We don’t need to maintain two pointers for front and rear if we use circular linked list. We can maintain a pointer to the last inserted node and front can always be obtained as next of last.
- Circular lists are useful in applications to repeatedly go around the list. For example, when multiple applications are running on a PC, it is common for the operating system to put the running applications on a list and then to cycle through them, giving each of them a slice of time to execute, and then making them wait while the CPU is given to another application. It is convenient for the operating system to use a circular list so that when it reaches the end of the list it can cycle around to the front of the list.
- 4) Circular Doubly Linked Lists are used for implementation of advanced data structures like Fibonacci Heap.

 */

type CircularLinkedList struct {
	head *Node
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{}
}

// Create Circular linked list, add element at the front
func (cll *CircularLinkedList) AddElementAtFront(value int) {
	var newNode = Node{data: value}

	tempHead := cll.head

	if cll.head == nil {
		cll.head = &newNode
		newNode.next = &newNode
		return
	}

	if tempHead.next == tempHead {
		newNode.next = tempHead
		tempHead.next = &newNode
		cll.head = &newNode
		return
	}

	for ; tempHead.next != cll.head; {
		tempHead = tempHead.next
	}
	cll.head = &newNode
	newNode.next = tempHead.next
	tempHead.next = &newNode
}

func (cll *CircularLinkedList) Print() {
	if cll.head == nil {
		return
	}

	log.Println(cll.head.data)

	if cll.head.next == cll.head {
		return
	}
	tempHead := cll.head
	for tempHead = tempHead.next; tempHead != cll.head; {
		print(tempHead.data)
		log.Println(tempHead.data)

		tempHead = tempHead.next
	}
	return
}

//TODO
// Circular Linked List Introduction and Applications,
// Circular Linked List Traversal
// Split a Circular Linked List into two halves
// Sorted insert for circular linked list
// Check if a linked list is Circular Linked List
// Convert a Binary Tree to a Circular Doubly Link List
// Circular Singly Linked List | Insertion
// Deletion from a Circular Linked List
// Circular Queue | Set 2 (Circular Linked List Implementation)
// Count nodes in Circular linked list
// Josephus Circle using circular linked list
// Convert singly linked list into circular linked list
// Circular Linked List | Set 1 (Introduction and Applications)
// Circular Linked List | Set 2 (Traversal)
// Implementation of Deque using circular array
// Exchange first and last nodes in Circular Linked List
