package pipeline

type Pipe interface {
	Handle(data interface{}, next func(data interface{}) interface{}) interface{}
}
