package http

type Handlers struct{}

type HandlerDeps struct{}

var H *Handlers

func InitializeHandlers(deps *HandlerDeps) {
	H = &Handlers{}
}
