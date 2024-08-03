package handler

type Middleware interface {
	Handle(Controller) Controller
}

type MiddlewareChain struct {
	Middlewares []Middleware
}

func (c MiddlewareChain) Handle(handler Controller) Controller {
	for _, mw := range c.Middlewares {
		handler = mw.Handle(handler)
	}

	return handler
}
