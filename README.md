Phalcon Go Framework
====================

Phalcon GO is a web for the GO language providing high performance and lower resource consumption.

Example
-------

```go

package main

import p "phalcon"

func main() {

    app := &p.Micro {}

    // Add handler to server's map.
    app.Map("/", func(c *p.AppContext) {
        c.GetResponse().SetContent("<h1>Welcome!</H1>!!")
    })

    // Add handler to server's map.
    app.Map("/hello", func(c *p.AppContext) {
        c.GetResponse().SetContent("<h1>Hello</H1>!!")
    })

    app.Map("/say/{name}/{surname}", func(c *p.AppContext) {
        c.GetResponse().SetContent("<h1>Hi " + c.GetParam("name") + ", " + c.GetParam("surname") + "!!</h1>")
    })

    app.Handle()
}

```

License
-------
Phalcon is open-sourced software licensed under the New BSD License. See the docs/LICENSE.txt file for more information.