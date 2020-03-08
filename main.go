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

	i := 1
	for i <= 10 {
		ll.AddElementAtFront(i)
		i++
	}

	ll.Print()
	for i <= 20 {
		ll.AddElementAtFront(i)
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