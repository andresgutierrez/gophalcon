package phalcon

import "fmt"
import "net/http"

type Micro struct {
    di *DI
    handlers [32]func(*AppContext)
    httpWriter http.ResponseWriter
}

func (this *Micro) Map(pattern string, handler func(*AppContext)) *Route {
    route := this.GetDI().GetRouter().Add(pattern)
    this.handlers[route.GetId()] = handler
    return route
}

func (this *Micro) SetDI(di *DI) *Micro {
    this.di = di
    return this
}

func (this *Micro) GetDI() *DI {
    if this.di == nil {
        this.di = &DI {};
    }
    return this.di
}

func (this *Micro) GetResponse() *Response {
    return this.GetDI().GetResponse()
}

func (this *Micro) ServeHTTP(httpWriter http.ResponseWriter, httpRequest *http.Request) {

    this.httpWriter = httpWriter

    fmt.Println(httpRequest.URL.Path)

    router := this.GetDI().GetRouter()

    match := router.Handle(httpRequest.URL.Path)

    if (match.WasMatched()) {

        //Create a context
        context := &AppContext {}

        request := &Request {HttpRequest: httpRequest}
        context.SetRequest(request)

        response := &Response {}
        context.SetResponse(response)

        route := match.GetMatchedRoute()
        context.SetParams(match.GetParams())

        this.handlers[route.GetId()](context)

        httpWriter.Write([]byte(context.GetResponse().GetContent()))
        return
    }

    handler := http.NotFoundHandler()
    handler.ServeHTTP(httpWriter, httpRequest)
}

func (this *Micro) Handle() {
    http.ListenAndServe(":3030", this)
}

