package binaryheap

import "errors"

type Heap struct {
	content []Node
}

type Node interface {
	Value() int64
}

func New(capacity int) *Heap {
	return &Heap{
		content: make([]Node, 0, capacity),
	}
}

func (heap *Heap) Add(node Node) {
	heap.content = append(heap.content, node)
	//heapify by bubbling up
	indexOfNew := len(heap.content) - 1
	for ;heap.content[indexOfNew].Value() < heap.content[heap.parent(indexOfNew)].Value() && indexOfNew != 0; {
		heap.swap(indexOfNew, heap.parent(indexOfNew))
		indexOfNew = heap.parent(indexOfNew)
	}
}
func (heap *Heap) RemoveTop() (node Node, err error) {
	length := len(heap.content)
	if length == 0 {
		return nil, errors.New("heap is empty")
	}
	node = heap.content[0]
	leaf := heap.content[length - 1]
	heap.content = heap.content[:length - 1]
	length -= 1
	if length == 0 {
		return node, nil
	}
	heap.content[0] = leaf
	//heapify by sinking the leaf back down

	for
	currentPosition, left, right := 0, heap.left(0),heap.right(0);
		right < length && heap.content[right].Value() < heap.content[currentPosition].Value() || left < length && heap.content[left].Value() < heap.content[currentPosition].Value();
	left, right =  heap.left(currentPosition),heap.right(currentPosition){
		//swap with smaller
		if left < length && right >= length ||  heap.content[left].Value() < heap.content[right].Value() {
			heap.swap(currentPosition, left)
			currentPosition = left
		} else if right < length {
			heap.swap(currentPosition, right)
			currentPosition = right
		}
	}
	return
}
func (heap *Heap) Top() (Node, error) {
	if len(heap.content) == 0 {
		return nil, errors.New("empty")
	}
	return heap.content[0], nil
}

func (heap *Heap) parent (index int) int {
	return (index-1) / 2
}
func (heap *Heap) left (index int) int {
	return (index*2) + 1
}
func (heap *Heap) right (index int) int {
	return heap.left(index) + 1
}
func (heap *Heap) swap(i,j int) []Node {
	tmp := heap.content[i]
	heap.content[i] = heap.content[j]
	heap.content[j] = tmp
	return heap.content
}


