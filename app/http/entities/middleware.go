package entities

type Middleware func(Controller) Controller

func MiddlewareChain(middlewares []Middleware) Middleware {
    return func(handler Controller) Controller {
        for _, mw := range middlewares {
            handler = mw(handler)
        }
        return handler
    }
}
