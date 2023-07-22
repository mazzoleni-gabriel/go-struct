package quicksort

type QuickSorter struct {
	data []int
}

func New(data []int) QuickSorter {
	return QuickSorter{data: data}
}

func (s *QuickSorter) Get() []int {
	return s.data
}

func (s *QuickSorter) Sort() {
	quickSort(0, len(s.data)-1, s.data)
}

func quickSort(start int, finish int, data []int) {
	if start >= finish {
		return
	}
	pivot := finish
	i := start
	j := start - 1

	for ; i <= finish; i++ {
		if data[i] >= data[pivot] {
			continue
		}

		j++
		temp := data[j]
		data[j] = data[i]
		data[i] = temp
	}

	//Swap pivot
	j++
	temp := data[j]
	data[j] = data[pivot]
	data[pivot] = temp

	middle := j
	quickSort(start, middle-1, data)
	quickSort(middle+1, finish, data)
}
