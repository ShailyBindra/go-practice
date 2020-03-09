package main

import (
	"fmt"
	"github.com/ShailyBindra/go-practice/linkedlist"
	"log"
)

func main() {
	log.Println("Starting")

	//test linked list
	testLinkedList()

	//test stack

}

func testLinkedList() {

	ll := linkedlist.NewLinkedList()

	i := 10
	for i > 0 {
		ll.AddElementAtFront(i)
		i--
	}
	ll.Print()

	for i = 11; i <= 20; {
		ll.AddElementAtEnd(i)
		i++
	}
	ll.Print()

	log.Println("m\n\nn")

	ll.DeleteAtGivenIndex(2)
	ll.Print()

	fmt.Print(ll.Length())

	log.Println("\n\n")

	ll.Reverse()
	ll.Print()

	ll.AddElementAtGivenIndex(2000, 20)
	ll.Print()

}
