// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

// Element is a structure containing keys and data values inside
// each element of PriorityQueue.
type Element struct {
	Key   int
	Value interface{}
	// Index tracks the position of the element inside the underlying array
	Index int
}

// PriorityQueue is a priority queue data implemented
// with a binary heap.
type PriorityQueue interface {
	// Insert one element to the queue, given element.Key
	Insert(element *Element)
	// Remove the first element off the queue, which would have the lowest element.Key value
	ExtractMin() *Element
	// Change the position (priority) in the queue if there's a change in
	// element.Key
	DecreaseKey(element *Element)
	// Get the current number of element remaining in the queue
	GetLength() int
}

// HeapSort is an implementation of heap sort algorithm.
func HeapSort(A []float64, n int) {

}
