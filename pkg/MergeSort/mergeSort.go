package MergeSort

//merge sort function
func MergeSort(unsortedSlice []int, startoffset int, endoffset int) {
	if startoffset < endoffset {
		midpoint := (startoffset + endoffset) / 2
		MergeSort(unsortedSlice, startoffset, midpoint)
		MergeSort(unsortedSlice, midpoint+1, endoffset)
		// Merge left and right partitions
		merge(unsortedSlice, startoffset, midpoint, endoffset)
	}

}


func merge(unsortedSlice []int, startoffset int, midpoint int, endoffset int) {
	s := endoffset - startoffset + 1
    mergedNumbers := make([]int,s)

	leftPos := startoffset
	rightPos := midpoint + 1
	mergePos := 0
	for leftPos <= midpoint && rightPos <= endoffset {
			if unsortedSlice[leftPos] <= unsortedSlice[rightPos] {
				mergedNumbers[mergePos] = unsortedSlice[leftPos]
				leftPos++
			} else {
				mergedNumbers[mergePos] = unsortedSlice[rightPos]
				rightPos++
			}
			mergePos++
		}

	// If left partition not empty, add remaining elements
	for leftPos <= midpoint {
		mergedNumbers[mergePos] = unsortedSlice[leftPos]
		leftPos++
		mergePos++
	}
	// If right partition not empty, add remaining elements
	for rightPos <= endoffset {
		mergedNumbers[mergePos] = unsortedSlice[rightPos]
		rightPos++
		mergePos++
	}

	for mergePos=0 ; mergePos < len(mergedNumbers); mergePos++ {
		unsortedSlice[startoffset+mergePos] = mergedNumbers[mergePos]
	}
}



