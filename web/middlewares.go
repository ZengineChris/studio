package web

import (
	"fmt"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/zenginechris/studio/domain"
)

// GetAuthMiddleware looks for a user in the session
// and bind it in the context
func GetAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := c.(*StudioContext)
			// look if the session key exist
			sess, err := session.Get("session", c)
			if err != nil {
				cc.User = domain.User{
					Role: domain.GuestUser,
				}
				// do not return an error
				return next(cc)
			}

			// token := sess.Values["access_token"]
			profile := sess.Values["profile"].(map[string]any)
			user := domain.User{
				ID:      fmt.Sprintf("%v", profile["sub"]),
				Role:    domain.MemberUser,
				Picture: fmt.Sprintf("%v", profile["picture"]),
				Name:    fmt.Sprintf("%v", profile["name"]),
			}
			// parse the user from the session

			cc.User = user
			return next(cc)
		}
	}
}
