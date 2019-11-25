package http

type ResponseStruct struct {
	content string
}

func Response() ResponseStruct {
	return ResponseStruct{}
}

func (r ResponseStruct) Content() string {
	return r.content
}

func (r ResponseStruct) SetContent(content string) ResponseStruct {
	r.content = content

	return r
}
