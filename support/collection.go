package support

type Items map[string]interface{}
type Item string

type CollectionStruct struct {
	items Items
}

func Collection(items Items) CollectionStruct {
	return CollectionStruct{items: items}
}
