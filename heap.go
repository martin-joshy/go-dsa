package dsa

type MinHeap struct {
	arr []int
}

func (h *MinHeap) HeapifyDown(index int) {
	lastIndex := len(h.arr) - 1

	for {
		left := index*2 + 1
		right := index*2 + 2
		smallest := index

		if right <= lastIndex && h.arr[left] < h.arr[smallest] {
			smallest = left
		}
		if left <= lastIndex && h.arr[right] < h.arr[smallest] {
			smallest = right
		}

		if smallest == index {
			break
		}

		h.arr[index], h.arr[smallest] = h.arr[smallest], h.arr[index]
		index = smallest

	}
}

func (h *MinHeap) HeapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2

		if h.arr[index] < h.arr[parent] {
			h.arr[index], h.arr[parent] = h.arr[parent], h.arr[index]
			index = parent
		} else {
			break
		}
	}
}

func (h *MinHeap) Insert(value int) {
	h.arr = append(h.arr, value)

	if len(h.arr) == 0 {
		return
	}

	lastIndex := len(h.arr) - 1
	h.HeapifyUp(lastIndex)
}

func (h *MinHeap) GetMin(index int) int {
	if len(h.arr) == 0 {
		return -1
	}

	min := h.arr[0]
	last := h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	if len(h.arr) > 0 {
		h.arr[0] = last
		h.HeapifyDown(0)
	}

	return min
}
