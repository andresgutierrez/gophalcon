package phalcon

import "fmt"
import "net/http"
import "os"
import "bufio"

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

    uri := httpRequest.URL.Path
    fmt.Println(uri)

    //Check if the file exists
    fi, err := os.Stat("public" + uri)
    if err == nil {
        if !fi.IsDir() {

            var (
                file *os.File
                part []byte
                prefix bool
            )

            if file, err = os.Open("public" + uri); err != nil {
                return
            }

            reader := bufio.NewReader(file)
            for {
                if part, prefix, err = reader.ReadLine(); err != nil {
                    break
                }
                fmt.Println(prefix)
                httpWriter.Write(part)
            }
            return
        }
    }

    //Try match a route
    router := this.GetDI().GetRouter()

    match := router.Handle(uri)

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
