# Pongo2Echo

[![CircleCI](https://circleci.com/gh/h3poteto/pongo2echo.svg?style=svg)](https://circleci.com/gh/h3poteto/pongo2echo)
[![GoDoc](https://godoc.org/github.com/h3poteto/pongo2echo?status.svg)](https://godoc.org/github.com/h3poteto/pongo2echo)

`pongo2echo` provides [pongo2](https://github.com/flosch/pongo2) renderer for [echo](https://github.com/labstack/echo) which is web application framework written by golang.
Pongo2 is a template engine likes django-syntax for golang.

This package is useful when you use pongo2 template in echo.

## Usage

Setup pongo2echo render when initialize echo.

```go

import (
	"github.com/h3poteto/pongo2echo"
	"github.com/labstack/echo"
)

func main() {
	render := pongo2echo.NewRenderer()
	render.AddDirectory("server/templates")
	e := echo.New()
	e.Renderer = render

	e.GET("/", func(c echo.Context) error {
		// index.html.tpl is located in server/templates/index.html.tpl
		return c.Render(http.StatusOK, "index.html.tpl",  map[string]interface{}{"title": "Index"})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

You can use pongo2 templte in `Render` function, and pass variables to template.

## License
The package is available as open source under the terms of the MIT License.
