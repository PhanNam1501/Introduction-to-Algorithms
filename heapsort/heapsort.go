package heapsort

import "log"

var maxInt = int(^uint(0) >> 1)
var minInt = -maxInt - 1

type MaxHeap struct {
	data     []int
	heapsize int
}

func NewMaxHeap(data []int) *MaxHeap {
	h := &MaxHeap{
		heapsize: len(data),
	}
	h.data = make([]int, len(data)+1)
	for i := 1; i < len(data)+1; i++ {
		h.data[i] = data[i-1]
	}
	return h
}

func (h *MaxHeap) GetMaxHeap() []int {
	return h.data[1:]
}

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func (h *MaxHeap) MaxHeapify(i int) {
	l := left(i)
	r := right(i)
	var largest int
	if l <= h.heapsize && h.data[l] > h.data[i] {
		largest = l
	} else {
		largest = i
	}

	if r <= h.heapsize && h.data[r] > h.data[largest] {
		largest = r
	}

	if largest != i {
		t := h.data[i]
		h.data[i] = h.data[largest]
		h.data[largest] = t
		h.MaxHeapify(largest)
	}
}

func (h *MaxHeap) BuildMaxHeap() {
	n := h.heapsize
	for i := n / 2; i > 0; i-- {
		h.MaxHeapify(i)
	}
}

func (h *MaxHeap) HeapSort() {
	h.BuildMaxHeap()
	n := h.heapsize
	for i := n; i > 1; i-- {
		t := h.data[1]
		h.data[1] = h.data[i]
		h.data[i] = t
		h.heapsize--
		h.MaxHeapify(1)
	}
	h.heapsize = n
}

func (h *MaxHeap) Maximum() int {
	if h.heapsize < 1 {
		log.Fatal("heap underflow")
	}
	return h.data[1]
}

func (h *MaxHeap) ExtractMax() int {
	max := h.Maximum()
	h.data[1] = h.data[h.heapsize]
	h.data = h.data[:h.heapsize]
	h.heapsize--
	h.MaxHeapify(1)
	return max
}

func (h *MaxHeap) IncreaseKey(x, key int) {
	if key < h.data[x] {
		log.Fatal("Invalid key")
	}
	h.data[x] = key
	for x > 1 && h.data[parent(x)] < h.data[x] {
		t := h.data[x]
		h.data[x] = h.data[parent(x)]
		h.data[parent(x)] = t
		x = parent(x)
	}
}

func (h *MaxHeap) Insert(x int) {
	h.heapsize++
	h.data = append(h.data, minInt)
	h.IncreaseKey(h.heapsize, x)
}
