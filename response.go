package phalcon

type Response struct {
    content string
}

func (this *Response) SetContent(content string) *Response {
    this.content = content
    return this
}

func (this *Response) GetContent() string {
    return this.content
}