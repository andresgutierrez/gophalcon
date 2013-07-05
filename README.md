Gooky Framework
====================

Gooky is a web framework for GO language providing high performance and lower resource consumption.

How to install?
-------
To get Gooky Framework, run this in your terminal:

go get github.com/gooky/gooky


Example
-------

```go

package main

import g "github.com/gooky/gooky"

func main() {

    app := &g.Micro {}

    // Add handler to server's map.
    app.Map("/", func(c *g.AppContext) {
        c.GetResponse().SetContent("<h1>Welcome!</H1>!!")
    })

    // Add handler to server's map.
    app.Map("/hello", func(c *g.AppContext) {
        c.GetResponse().SetContent("<h1>Hello</H1>!!")
    })

    app.Map("/say/{name}/{surname}", func(c *g.AppContext) {
        c.GetResponse().SetContent("<h1>Hi " + c.GetParam("name") + ", " + c.GetParam("surname") + "!!</h1>")
    })

    app.Handle()
}

```

License
-------
Gooky Framework is open-sourced software licensed under the New BSD License. See the docs/LICENSE.txt file for more information.