package sort

import "fmt"

type minheap struct {
	arr []int
}

func newMinHeap(arr []int) *minheap {
	minheap := &minheap{
		arr: arr,
	}
	return minheap
}

func (m *minheap) leftChildIndex(index int) int {
	return 2*index + 1
}

func (m *minheap) rightChildIndex(index int) int {
	return 2*index + 2
}

func (m *minheap) swap(first, second int) {
	temp := m.arr[first]
	m.arr[first] = m.arr[second]
	m.arr[second] = temp
}

func (m *minheap) isLeaf(index int, size int) bool {
	if index >= (size/2) && index <= size {
		return true
	}
	return false
}

func (m *minheap) downHeapify(current int, size int) {
	if m.isLeaf(current, size) {
		return
	}

	smallest := current
	leftChildIndex := m.leftChildIndex(current)
	rightChildIndex := m.rightChildIndex(current)

	if leftChildIndex < size && m.arr[leftChildIndex] < m.arr[smallest] {
		smallest = leftChildIndex
	}

	if rightChildIndex < size && m.arr[rightChildIndex] < m.arr[smallest] {
		smallest = rightChildIndex
	}

	if smallest != current {
		m.swap(current, smallest)
		m.downHeapify(smallest, size)
	}
}

func (m *minheap) buildMinHeap(size int) {
	for index := ((size / 2) - 1); index >= 0; index-- {
		m.downHeapify(index, size)
	}
}

func (m *minheap) sort(size int) {
	m.buildMinHeap(size)
	for i := size - 1; i > 0; i-- {
		m.swap(0, i)
		m.downHeapify(0, i)
	}
}

func (m *minheap) reverse() {
	for i, j := 0, len(m.arr)-1; i < j; i, j = i+1, j-1 {
		m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
	}
}

func (m *minheap) print() {
	for _, val := range m.arr {
		fmt.Println(val)
	}
}

func TestMinHeapSort() {
	inputArray := []int{4, 8, 22, 0, -2, 70, 19}
	minheap := newMinHeap(inputArray)
	minheap.sort(len(inputArray))
	minheap.print()
	minheap.reverse()
	minheap.print()
}
