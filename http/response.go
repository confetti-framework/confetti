package http

type Response struct {
	content string
}

func NewResponse() Response {
	return Response{}
}

func (r Response) Content() string {
	return r.content
}

func (r Response) SetContent(content string) Response {
	r.content = content

	return r
}
