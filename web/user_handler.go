package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zenginechris/studio/infrastructure"
	"github.com/zenginechris/studio/views/pages"
)

func UserHomeHandler(repo infrastructure.CalendarRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*StudioContext)

		classes, err :=  repo.GetEvents(c.Request().Context())
		if err != nil {
			return err 
		}

		vm := pages.UserHomeViewModel{
			User: cc.User,
			Classes: classes,
		}
		return Render(c, http.StatusOK, pages.UserHome(vm))
	}
}
