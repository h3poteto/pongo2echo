package pongo2echo

import (
	"bytes"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func TestInterface(t *testing.T) {
	renderer := NewRenderer()
	var _ echo.Renderer = renderer
}

func TestRender(t *testing.T) {
	renderer := NewRenderer()
	renderer.AddDirectory("./templates")

	buf := new(bytes.Buffer)
	err := renderer.Render(buf, "test.html", map[string]interface{}{"name": "h3poteto"}, nil)
	if err != nil {
		t.Error("Cloud not render test.html through pongo2")
	}

	contains := strings.Contains(buf.String(), "h3poteto")
	if contains != true {
		t.Errorf("Rendered html does not contain name parameter: %s", buf.String())
	}
}
