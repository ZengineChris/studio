package web

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/zenginechris/studio/config"
	"github.com/zenginechris/studio/domain"
)

type StudioContext struct {
	User domain.User
	echo.Context
}

func NewServer(config *config.Config) *echo.Echo {
	e := echo.New()
	// build the custome context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &StudioContext{
				domain.User{},
				c,
			}
			return next(cc)
		}
	})
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	// use the user middleware 
	e.Static("/assets", "assets")
	// registere the Middleware
	// register the routes
	RegisterRoutes(e)
	return e
}
