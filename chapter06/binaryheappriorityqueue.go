// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

// NewBinaryHeapPriorityQueue returns a new instance of PriorityQueue,
// which is actually an interface and implemented with
// binaryHeapPriorityQueue struct. The reason why we are
// not directly exporting binaryHeapPriorityQueue is because
// we want to enforce Length: 0 when creating an instance of
// PriorityQueue.
func NewBinaryHeapPriorityQueue() PriorityQueue {
	return &binaryHeapPriorityQueue{
		Data: make([]*Element, 0),
	}
}

type binaryHeapPriorityQueue struct {
	Data []*Element
}

func (Q *binaryHeapPriorityQueue) Insert(element *Element) {

	element.Index = len(Q.Data)
	Q.Data = append(Q.Data, element)

	// If the queue was empty, we just add key as a first element
	if len(Q.Data) > 1 {
		Q.satisfyHeapProperty((len(Q.Data) - 2) / 2)
	}
}

func (Q *binaryHeapPriorityQueue) ExtractMin() *Element {

	if len(Q.Data) == 0 {
		return nil
	}

	root := make([]*Element, 1)
	copy(root, Q.Data[0:1])

	if len(Q.Data) == 1 {
		Q.Data[0] = nil
		Q.Data = Q.Data[:0]
	} else {
		copy(Q.Data[0:], Q.Data[1:])
		Q.Data[len(Q.Data)-1] = nil
		Q.Data = Q.Data[:len(Q.Data)-1]

		newRootKey := Q.Data[0].Key
		newRootValue := Q.Data[0].Value

		Q.Data[0].Key = Q.Data[len(Q.Data)-1].Key
		Q.Data[0].Value = Q.Data[len(Q.Data)-1].Value

		Q.Data[len(Q.Data)-1].Key = newRootKey
		Q.Data[len(Q.Data)-1].Value = newRootValue
		// If there was at least one existing element,
		// we need to check for heap property and swap elements if needed
		Q.satisfyHeapProperty(0)
	}

	return root[0]
}

func (Q *binaryHeapPriorityQueue) satisfyHeapProperty(parentIndex int) {
	heapPropertySatisfied := false

	leftChildIndex := 2*parentIndex + 1
	rightChildIndex := leftChildIndex + 1

	for !heapPropertySatisfied {
		// if parentIndex >= len(Q.Data), it is asking for
		// non-existent node. We can terminate the loop.
		if parentIndex >= len(Q.Data) {
			heapPropertySatisfied = true
		} else {
			if leftChildIndex < len(Q.Data) && Q.Data[parentIndex].Key > Q.Data[leftChildIndex].Key {
				// If the left child exists & its key is smaller than the parent's,
				// then swap the two.
				leftChildKey := Q.Data[leftChildIndex].Key
				leftChildValue := Q.Data[leftChildIndex].Value

				Q.Data[leftChildIndex].Key = Q.Data[parentIndex].Key
				Q.Data[leftChildIndex].Value = Q.Data[parentIndex].Value

				Q.Data[parentIndex].Key = leftChildKey
				Q.Data[parentIndex].Value = leftChildValue
			} else if rightChildIndex < len(Q.Data) && Q.Data[parentIndex].Key > Q.Data[rightChildIndex].Key {
				// If the right child exists & its key is smaller than the parent's,
				// then swap the two. We can safely assume that there's no case where
				// the left child doesn't exist but the right does, since Q is a binary heap.
				rightChildKey := Q.Data[rightChildIndex].Key
				rightChildValue := Q.Data[rightChildIndex].Value

				Q.Data[rightChildIndex].Key = Q.Data[parentIndex].Key
				Q.Data[rightChildIndex].Value = Q.Data[parentIndex].Value

				Q.Data[parentIndex].Key = rightChildKey
				Q.Data[parentIndex].Value = rightChildValue
			} else {
				// if the heap property seems to have been met for parentIndex
				// and its children, then we also need to make sure the property holds
				// for each child and their descedants (grandchildren) as well.
				Q.satisfyHeapProperty(leftChildIndex)
				Q.satisfyHeapProperty(rightChildIndex)

				// After checks above are complete, we can now conclude that heap property
				// holds for Q.
				heapPropertySatisfied = true
			}
		}
	}
}

func (Q *binaryHeapPriorityQueue) DecreaseKey(element *Element) {
	Q.satisfyHeapProperty(element.Index)
}

func (Q *binaryHeapPriorityQueue) GetLength() int {
	return len(Q.Data)
}