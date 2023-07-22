package mergesort

type Sorter struct {
	data []int
}

func New(data []int) Sorter {
	return Sorter{data: data}
}

func (s *Sorter) Get() []int {
	return s.data
}

func (s *Sorter) Sort() {
	s.data = mergeSort(s.data)
}

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	leftArray, rightArray := split(data)

	leftArray = mergeSort(leftArray)
	rightArray = mergeSort(rightArray)
	return merge(leftArray, rightArray, data)
}

func merge(leftArray []int, rightArray []int, data []int) []int {
	leftIdx := 0
	lenLeft := len(leftArray)
	rightIdx := 0
	lenRight := len(rightArray)
	newData := make([]int, len(data))

	for dataIdx, _ := range data {
		// Left array is finished, only add remaining right array elements
		if leftIdx >= lenLeft {
			newData[dataIdx] = rightArray[rightIdx]
			rightIdx++
			continue
		}
		// Right array is finished, only add remaining left array elements
		if rightIdx >= lenRight {
			newData[dataIdx] = leftArray[leftIdx]
			leftIdx++
			continue
		}

		if leftArray[leftIdx] < rightArray[rightIdx] {
			newData[dataIdx] = leftArray[leftIdx]
			leftIdx++
			continue
		}
		if rightArray[rightIdx] < leftArray[leftIdx] {
			newData[dataIdx] = rightArray[rightIdx]
			rightIdx++
			continue
		}

	}

	return newData
}

func split(data []int) (left []int, right []int) {
	middle := len(data) / 2
	left = data[:middle]
	right = data[middle:]
	return
}
