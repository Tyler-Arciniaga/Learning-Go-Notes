package main

import (
	"errors"
	"fmt"
)

type MinHeap[T any] struct{
	slice []T
	comparator func(a, b T) bool //comparator that returns true if a is comparator than b
}

func (m *MinHeap[T]) bubble_up(i int) {
	curr := i
	for curr != 0 && m.comparator(m.slice[curr], m.slice[curr / 2]){
		temp := m.slice[curr / 2]
		m.slice[curr / 2] = m.slice[curr]
		m.slice[curr] = temp
		curr = curr / 2
	}
}

func (m *MinHeap[T]) bubble_down(i int){
	curr := i

	for (curr * 2) < len(m.slice){
		temp := m.slice[curr]
		leftChild := m.slice[curr * 2]
		rightChild := temp
		if ((curr * 2) + 1) < len(m.slice){
			rightChild = m.slice[(curr * 2) + 1]
		}
		if m.comparator(leftChild, temp) || m.comparator(rightChild, temp){
			if m.comparator(leftChild, rightChild){
				m.slice[curr] = leftChild
				m.slice[curr * 2] = temp
				curr = curr * 2
			} else {
				m.slice[curr] = rightChild
				m.slice[(curr * 2) + 1] = temp
				curr = (curr * 2) + 1
			}
			continue
		}
		break
	}
}

func (m *MinHeap[T]) Push(x T){
	m.slice = append(m.slice, x)
	m.bubble_up(len(m.slice) - 1)
}

func (m *MinHeap[T]) Peek() T{
	return m.slice[0]
}

func (m *MinHeap[T]) Pop() (T, error){
	if len(m.slice) == 0{
		var zero T
		return zero, errors.New("cannot pop off a empty heap")
	}
	top_element := m.slice[0]
	m.slice[0] = m.slice[len(m.slice) - 1]
	m.slice = m.slice[:len(m.slice) - 1]
	m.bubble_down(0)
	return top_element, nil
}

func (m *MinHeap[T]) Len() int {
	return len(m.slice)
}

func main(){
	Test1()
	Test2()
	Test3()
	Test4()
}

//p.s. Transfer to Go Testing eventually, using function for now
func Test1(){
	fmt.Println("TEST 1")
	h := MinHeap[int]{slice: []int{}, comparator: func(a, b int) bool {
		return a < b
	}}
	h.Push(5)
	h.Push(7)
	h.Push(1)
	h.Push(2)
	fmt.Printf("Length: %d\n", h.Len())
	x, _ := h.Pop()
	fmt.Printf("Popped: %d\n", x)

	fmt.Printf("Top: %d\n", h.Peek())
	h.Pop()
	fmt.Printf("Top: %d\n", h.Peek())
	h.Push(6)
	h.Push(0)
	x, _ = h.Pop()
	fmt.Printf("Popped: %d\n", x)
}

func Test2(){
	fmt.Println("\nTEST 2")
	h := MinHeap[int]{slice: []int{}, comparator: func(a, b int) bool {
		return a < b
	}}
	input := []int{5, 4, 7, 0, 1, -1, 4, 2, 8, 3}
	for _, v := range input{
		h.Push(v)
	}

	for h.Len() > 0{
		x, _ := h.Pop()
		fmt.Println(x)
	}
}

func Test3(){
	fmt.Println("\nTEST 3")
	h := MinHeap[string]{slice: []string{}, comparator: func(a, b string) bool {
		return len(a) < len(b)
	}}
	input := []string{"Tyler", "Allyson", "Vicky", "Ryan", "Ori", "zzzzzzzzzz", "x"}
	for _, v := range input{
		h.Push(v)
	}

	for h.Len() > 0{
		x, _ := h.Pop()
		fmt.Println(x)
	}
}

// max heap test
func Test4(){
	fmt.Println("\nTEST 2")
	h := MinHeap[int]{slice: []int{}, comparator: func(a, b int) bool {
		return a > b
	}}
	input := []int{5, 4, 7, 0, 1, -1, 4, 2, 8, 3}
	for _, v := range input{
		h.Push(v)
	}

	for h.Len() > 0{
		x, _ := h.Pop()
		fmt.Println(x)
	}
}