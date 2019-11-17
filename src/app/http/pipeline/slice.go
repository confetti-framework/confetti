package pipeline

import (
	"lanvard/support/caller"
)

type Item PipeInterface
type Items []PipeInterface

type SliceStruct struct {
	Items Items
}

func Slice(items Items) SliceStruct {
	return SliceStruct{Items: items}
}

func (c SliceStruct) Push(item Item) SliceStruct {
	c.Items = append(c.Items, item)

	return c
}

func (c SliceStruct) Reverse() SliceStruct {
	items := c.Items
	for left, right := 0, len(items)-1; left < right; left, right = left+1, right-1 {
		items[left], items[right] = items[right], items[left]
	}

	c.Items = items

	return c
}

func (c SliceStruct) Path() string { return caller.Path() }

func (c SliceStruct) ToSlice() Items {
	return c.Items
}
