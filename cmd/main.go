package main

import (
	"encoding/gob"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/zenginechris/studio/config"
	"github.com/zenginechris/studio/web"
)

func main() {
	gob.Register(map[string]interface{}{})
	config := config.New()
	e := web.NewServer(config)
	e.Logger.Fatal(e.Start(":1323"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
