package main

import "fmt"

type list struct {
	data int
	next *list
}

type pointerToList struct {
	head *list
}

func initList() *pointerToList {
	return &pointerToList{}
}

func (l *pointerToList) addNode(data int) {
	node := &list{
		data: data,
	}
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
	return
}

func (l *pointerToList) printList() {
	node := l.head
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
	return
}

func (l *pointerToList) reverseList() *pointerToList {
	reversedList := initList()
	node := l.head
	for node != nil {
		reversedList.addNode(node.data)
		node = node.next
	}
	return reversedList
}

func main() {
	nodes := []int{
		1,
		2,
		3,
	}
	currentList := initList()
	for _, nodeData := range nodes {
		currentList.addNode(nodeData)
	}
	currentList.printList()
	reversedList := currentList.reverseList()
	reversedList.printList()
}
