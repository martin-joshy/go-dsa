package dsa

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		current := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = current

	}
}

func selectSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		smallest := i
		for j := i + 1; j <= n-1; j++ {
			if arr[smallest] > arr[j] {
				smallest = j
			}
		}

		arr[smallest], arr[i] = arr[i], arr[smallest]
	}
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return mergeSorted(left, right)
}

func mergeSorted(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
