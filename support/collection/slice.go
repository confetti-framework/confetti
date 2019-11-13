package collection

type Slice []interface{}

func (s Slice) Push(item interface{}) Slice {
	slice := append(s, item)

	return slice
}

func (s Slice) Reverse() Slice {
	slice := s
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}

	return slice
}
