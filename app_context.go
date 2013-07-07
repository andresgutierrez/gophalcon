package gooky

type AppContext struct {
	app *Micro
	response *Response
	request *Request
	params map[string]string
}

func (this *AppContext) SetRequest(request *Request) *AppContext {
    this.request = request
    return this
}

func (this *AppContext) SetResponse(response *Response) *AppContext {
    this.response = response
    return this
}

func (this *AppContext) SetParams(params map[string]string) *AppContext {
    this.params = params
    return this
}

func (this *AppContext) GetResponse() *Response {
    return this.response
}

func (this *AppContext) GetParam(param string) string {
    return this.params[param]
}
