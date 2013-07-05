package phalcon

type DI struct {
    router *Router
    response *Response
}

func (this *DI) GetRouter() *Router {
    if this.router == nil {
        this.router = &Router{}
    }
    return this.router
}

func (this *DI) GetResponse() *Response {
    if this.response == nil {
        this.response = &Response{}
    }
    return this.response
}