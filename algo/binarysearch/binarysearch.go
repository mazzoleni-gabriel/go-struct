package binarysearch

type BinarySearcher struct {
	data []int
}

func New(data []int) BinarySearcher {
	return BinarySearcher{data: data}
}

func (b BinarySearcher) Search(target int) bool {
	return search(b.data, target)
}

func search(data []int, target int) bool {
	if len(data) == 1 {
		return data[0] == target
	}
	middle := len(data) / 2
	if data[middle] == target {
		return true
	}
	//Search left half
	if target < data[middle] {
		return search(data[:middle], target)
	}

	//Search right half
	if target > data[middle] {
		return search(data[middle:], target)
	}
	return false
}
