package dsa

func heapSort(arr []int) {
	n := len(arr)

	for i := (n / 2) - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, lgt int, index int) {
	for {
		left := index*2 + 1
		right := index*2 + 2
		largest := index

		if left < lgt && arr[left] > arr[largest] {
			largest = left
		}
		if right < lgt && arr[right] > arr[largest] {
			largest = right
		}

		if largest != index {
			arr[index], arr[largest] = arr[largest], arr[index]
			index = largest
		} else {
			break
		}
	}
}
