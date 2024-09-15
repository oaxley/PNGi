package linkedlist

import "testing"

func TestAddOne(t *testing.T) {
	var list LinkedList

	// add one value
	list.Add(1)

	// check that both head and tail are equivalent
	if list.head != list.tail {
		t.Fatal("Head and Tail are not equal when list contains only 1 element")
	}
}

func TestAddTwo(t *testing.T) {
	var list LinkedList

	// add values
	list.Add(1)
	list.Add(2)

	// check head & tail are different
	if list.head == list.tail {
		t.Fatal("Head == Tail with 2 elements in the list")
	}

	// check head.next -> tail
	if list.head.Next != list.tail {
		t.Fatal("The two elements in the list are not linked (Head.Next != Tail)")
	}

	// check tail.next = nil
	if list.tail.Next != nil {
		t.Fatal("Tail *next does not point to nil")
	}
}

func TestCount(t *testing.T) {
	var list LinkedList

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	if list.count != 10 {
		t.Fatal("Count is not equal to 10")
	}
}

func TestGet(t *testing.T) {
	var list LinkedList

	for i := 0; i < 10; i++ {
		list.Add(i + 1)
	}

	// retrieve the element #7, which should be equal to 8
	index := 7
	expected := 8
	value, _ := list.Get(index)
	if value != expected {
		t.Fatalf("Get(%d) does not return the correct value: expected = %d, received = %d", index, expected, value)
	}
}

func TestElements(t *testing.T) {
	var list LinkedList
	max_value := 10

	for i := 0; i < max_value; i++ {
		list.Add(i + 1)
	}

	expected_sum := (max_value * (max_value + 1)) / 2
	sum := 0
	for node := range list.Elements() {
		sum = sum + node.Data.(int)
	}

	if sum != expected_sum {
		t.Fatalf("Elements() does not return the correct values: expected = %d, received = %d", expected_sum, sum)
	}
}
