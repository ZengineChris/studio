// Package web
package web

import (
	"github.com/labstack/echo/v4"
	"github.com/zenginechris/studio/infrastructure"
)

func RegisterRoutes(e *echo.Echo) {
	classRepo := infrastructure.NewClassRepository()
	registerRepo := infrastructure.NewRegistrationRepository()
	calendarRepo := infrastructure.NewCalendarRepository(classRepo, registerRepo)

	e.GET("/", UserHomeHandler(calendarRepo), GetAuthMiddleware())
	e.POST("/calendar/events/:id/register", RegisterToCarendarEventHandler(
		registerRepo, calendarRepo), GetAuthMiddleware())
	e.POST("/calendar/events/:id/unregister", RegisterToCarendarEventHandler(
		registerRepo, calendarRepo), GetAuthMiddleware())
}
