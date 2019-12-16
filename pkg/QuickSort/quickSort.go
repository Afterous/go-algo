package QuickSort



func Quicksort(l []int, i int, k int) int {
	j := 0
	if i >= k {
		return 0
	}

	/* Partition the array.
	   Value j is the location of last
	   element in low partition. */
	j = partition(l, i, k)
	/* Recursively sort low and high
	   partitions */
	Quicksort(l, i, j)
	Quicksort(l, j+1, k)
	return 0
}

func partition(l []int, i int, k int) int {
	done := false
	midpoint := i + (k-i)/2 //Pick middle value as pivot
	pivot := l[midpoint]
	low := i
	high := k

	for !done {
		for l[low] < pivot {
			low = low + 1
		} /* Increment low while l[l] < pivot */
		for pivot < l[high] {
			high = high - 1
		} /* Decrement high while pivot < l[h] */
		if low >= high {
			done = true
		} else {
			l[low], l[high] = l[high], l[low]
			low = low + 1
			high = high - 1
		}
	}
	return high
}