package linkedlist

import "errors"

type (
	Node struct {
		Data interface{}
		Next *Node
	}

	LinkedList struct {
		head  *Node
		tail  *Node
		count int
	}
)

// add a new element to the list
func (list *LinkedList) Add(value interface{}) {
	// create the new element
	node := Node{
		Data: value,
		Next: nil,
	}

	// increment the count
	list.count = list.count + 1

	// first element in the list
	if list.head == nil {
		list.head = &node
		list.tail = &node
		return
	}

	// add the element to the end of the list
	list.tail.Next = &node
	list.tail = &node
}

// return the number of element in the list
func (list *LinkedList) Count() int {
	return list.count
}

// return the element at the index
func (list *LinkedList) Get(index int) (interface{}, error) {
	if index >= list.count {
		return nil, errors.New("index greater than number of available elements")
	}

	count := 0
	current := list.head
	for count < index {
		current = current.Next
		count = count + 1
	}

	return current.Data, nil
}

// generator to return all the elements
func (list *LinkedList) Elements() <-chan Node {
	ch := make(chan Node)
	ptr := list.head
	go func() {
		defer close(ch)

		for ptr != nil {
			ch <- *ptr
			ptr = ptr.Next
		}
	}()

	return ch
}
