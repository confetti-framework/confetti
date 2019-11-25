package support

import "lanvard/support/caller"

type Item []interface{}
type Items []Item

type Slice struct {
	Items Items
}

func NewSlice(items Items) Slice {
	return Slice{Items: items}
}

func (c Slice) Push(item Item) Slice {
	c.Items = append(c.Items, item)

	return c
}

func (c Slice) Reverse() Slice {
	items := c.Items
	for left, right := 0, len(items)-1; left < right; left, right = left+1, right-1 {
		items[left], items[right] = items[right], items[left]
	}

	c.Items = items

	return c
}

func (c Slice) ToSlice() Items {
	return c.Items
}

func (c Slice) Path() string { return caller.Path() }
