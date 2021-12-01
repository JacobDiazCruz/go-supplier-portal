package main

import "fmt"

type MyHeap struct {
	array []int
}

func (h *MyHeap) Insert(i int) {
	h.array = append(h.array, i)
	h.maxHeapifyUp(len(h.array) - 1)
}

// maxHeapifyDown will heapify from bottom to top
func (h *MyHeap) maxHeapifyUp(index int) {
	for _, v := range h.array {
		fmt.Println(v)
		if h.array[parent(index)] < h.array[index] {
			h.swap(parent(index), index)
		}
		// index = parent(index)
	}
}

func (h *MyHeap) sort() {
	temp := 0
	for i := 0; i < len(h.array); i++ {
		for j := i; j < len(h.array); j++ {
			if h.array[i] > h.array[j] {
				temp = h.array[j]
				h.array[j] = h.array[i]
				h.array[i] = temp
			}
		}
		fmt.Println("hrhr")
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i + 1*2
}

func right(i int) int {
	return i + 2*2
}

func (h *MyHeap) swap(i1 int, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() &MyHeap{
	m := &MyHeap{
		array: []int{10, 20, 30, 5},
	}
	// err := []int{10, 20, 30, 5}
	// if err != nil {
	// 	fmt.Println("hert")
	// }
	m.sort()
	fmt.Println()
	// for _, v := range buildHeap {
	// 	m.Insert(v)
	// 	fmt.Println(m)
	// }
}
