package main

import (
	"fmt"
)

type List struct {
	head, tail *Item
    len  int
}

type Item struct {
	next, prev *Item
	value interface{}
}

// Gets the first value from the list
func (l *List) First() interface{} {
	if l.len != 0 {
		return l.head.Value()
	}
	return nil
}

// Gets the last value from the list
func (l *List) Last() interface{} {
	if l.len != 0 {
		return l.tail.Value()
	}
	return nil
}

// Gets the list's lenght
func (l *List) Len() int {
	return l.len
}

// Adds an item to the top of the list
func (l *List) PushFront(v interface{}) {
	newitem := &Item{value: v, next: l.head}
	if l.len == 0 {
		l.head = newitem
		l.tail = newitem
	} else {
		l.head.prev = newitem
		l.head = newitem
	}
	l.len++
}

// Adds an item at the end of the list
func (l *List) PushBack(v interface{}) {
	newitem := &Item{value: v, prev: l.tail}
	if l.len == 0 {
		l.head = newitem
		l.tail = newitem
	} else {
		l.tail.next = newitem
		l.tail = newitem
	}
	l.len++
 	
}

// Removes an item from the list by index
func (l *List) Remove(index int) {
	if l.len == 0 {
		fmt.Println("list is empty")
		return
	}
	if !(index >= 0 && index < l.len) {
		return
	}
	var item *Item

	item = l.head
	for i := 0; i != index; i, item = i+1, item.Next() {
	}
	
	if l.head == item {
		l.head = item.Next()
	}
	if l.tail == item {
		l.tail = item.Prev()
	}
	if item.prev != nil {
		item.prev.next = item.Next()
	}
	if item.next != nil {
		item.next.prev = item.Prev()
	}
	
	l.len--
}

// Return the value of the item
func (i *Item) Value() interface{} {
	return i.value
}

// Return the next item
func (i *Item) Next() *Item {
	return i.next
}

// Return the previous item
func (i *Item) Prev() *Item {
	return i.prev
}




func main() {
	list := &List{}
	list.PushFront("a")
	list.PushBack("b")
	list.PushBack("c")
	list.PushFront("d")
	

	fmt.Printf("%s\t\t%2d\n", "Длина", list.Len())
	fmt.Printf("%s\t%2s\n", "First item:", list.First())
	fmt.Printf("%s\t%2s\n", "Last item:", list.Last())

	// Remove first item. Now the first item should be "a"
	list.Remove(0)
	fmt.Println("------------------")

	fmt.Printf("%s\t\t%2d\n", "Длина", list.Len())
	fmt.Printf("%s\t%2s\n", "First item:", list.First())
	fmt.Printf("%s\t%2s\n", "Last item:", list.Last())
	
}