package main

import (
	"fmt"
	"github.com/ShailyBindra/go-practice/linkedlist"
	"log"
)

func main() {
	log.Println("Starting")

	//test linked list
	//testLinkedList()

	//test circular ll
	testCircularLinkedList()

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

	ll.PrintRecursively()
	ll.HasLoop()
	ll.SearchElement(6)
	ll.SearchElement(60)
	ll.SearchElement(2000)

}

func testCircularLinkedList() {
	cll := linkedlist.NewCircularLinkedList()
	cll.AddElementAtFront(5)
	cll.Print()
	log.Println("m\n\nn")

	cll.AddElementAtFront(6)
	cll.Print()
	log.Println("m\n\nn")

	cll.AddElementAtFront(7)
	cll.Print()
	log.Println("m\n\nn")

	cll.AddElementAtFront(8)
	cll.Print()
	log.Println("m\n\nn")

}