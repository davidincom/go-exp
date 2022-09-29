package containers

import "fmt"

type minheap struct {
	heapArray []int
	size      int
	maxsize   int
}

func newMinHeap(maxsize int) *minheap {
	fmt.Println("create new min-heap.")
	minheap := &minheap{
		heapArray: []int{},
		size:      0,
		maxsize:   maxsize,
	}
	return minheap

}

func (m *minheap) isLeaf(index int) bool {
	if index >= (m.size/2) && index <= m.size {
		return true
	}
	return false
}

func (m *minheap) parentIndex(index int) int {
	return (index - 1) / 2
}

func (m *minheap) leftChildIndex(index int) int {
	return 2*index + 1
}

func (m *minheap) rightChildIndex(index int) int {
	return 2*index + 2
}

func (m *minheap) insert(item int) error {
	if m.size >= m.maxsize {
		return fmt.Errorf("heap is full")
	}

	m.heapArray = append(m.heapArray, item)
	m.size++
	m.upHeapify(m.size - 1)
	return nil
}

func (m *minheap) swap(first, second int) {
	temp := m.heapArray[first]
	m.heapArray[first] = m.heapArray[second]
	m.heapArray[second] = temp
}

func (m *minheap) upHeapify(index int) {
	for m.heapArray[index] < m.heapArray[m.parentIndex(index)] {
		m.swap(index, m.parentIndex(index))
		index = m.parentIndex(index)
	}
}

func (m *minheap) downHeapify(current int) {
	if m.isLeaf(current) {
		return
	}

	smallest := current
	leftChildIndex := m.leftChildIndex(current)
	rightChildIndex := m.rightChildIndex(current)

	if leftChildIndex < m.size && m.heapArray[leftChildIndex] < m.heapArray[smallest] {
		smallest = leftChildIndex
	}

	if rightChildIndex < m.size && m.heapArray[rightChildIndex] < m.heapArray[smallest] {
		smallest = rightChildIndex
	}

	if smallest != current {
		m.swap(current, smallest)
		m.downHeapify(smallest)
	}
}

func (m *minheap) buildMinHeap() {
	for index := ((m.size / 2) - 1); index >= 0; index-- {
		m.downHeapify(index)
	}
}

func (m *minheap) remove() int {
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size-1]
	m.heapArray = m.heapArray[:(m.size)-1]
	m.size--
	m.downHeapify(0)
	return top
}

func TestMinHeap() {
	inputArray := []int{6, 5, 3, 7, 2, 8, 99, -1, 789, 20, 51}
	minHeap := newMinHeap(len(inputArray))
	for i := 0; i < len(inputArray); i++ {
		minHeap.insert(inputArray[i])
	}
	minHeap.buildMinHeap()
	for i := 0; i < len(inputArray); i++ {
		fmt.Println(minHeap.remove())
	}
	// fmt.Scanln()
}
