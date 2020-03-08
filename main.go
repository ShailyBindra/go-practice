package main

import (
	"github.com/ShailyBindra/go-practice/linkedlist"
	"log"
)

func main()  {
	log.Println("Starting")

	//test linked list
	testLinkedList()

	//test stack

}

func testLinkedList()  {

	ll := linkedlist.NewLinkedList()

	i := 10
	for i > 0 {
		ll.AddElementAtFront(i)
		i--
	}

	for i=11 ;i <= 20; {
		ll.AddElementAtEnd(i)
		i++
	}
	log.Println("n\n")

	ll.Print()

	//i = 11
	//for i <= 20 {
	//	ll.AddElementAtFront(i)
	//	i++
	//}
	//ll.Print()
}