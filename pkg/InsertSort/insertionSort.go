package InsertSort

func InsertionSort(l []int) {
	for i := 0; i < len(l)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(l); j++ {
			if l[minIndex] > l[j] {
				minIndex = j
			}
		}
		temp := l[i]
		l[i] = l[minIndex]
		l[minIndex] = temp
	}
}
