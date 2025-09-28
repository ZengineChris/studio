package web

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zenginechris/studio/domain"
	"github.com/zenginechris/studio/infrastructure"
	"github.com/zenginechris/studio/views/components"
)

// returns the actions component for a calendar event
func RegisterToCarendarEventHandler(
	registrationRepo infrastructure.RegistrationRepository,
	calendarRepo infrastructure.CalendarRepository,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*StudioContext)
		eventID := c.Param("id")
		cEvent := domain.CalendarEventID{}
		cEvent.FromString(eventID)
		fmt.Println(cEvent)
		event, err := calendarRepo.FindEvent(c.Request().Context(), cEvent)
		if err != nil {
			return err
		}
		registration := domain.Registration{
			ID:          "the",
			ReferenceID: event.ID(),
			UserID:      cc.User.ID,
		}

		_, err = registrationRepo.Store(c.Request().Context(), registration)
		if err != nil {
			return err
		}

		// here we search for the event again to get all registerd ussers

		return Render(c, http.StatusOK, components.CalendarEventActions(event))
	}
}

func UnRegisterToCarendarEventHandler(
	registrationRepo infrastructure.RegistrationRepository,
	calendarRepo infrastructure.CalendarRepository,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		// try to register to the event
		eventID := c.Param("id")
		cEvent := domain.CalendarEventID{}
		cEvent.FromString(eventID)
		event, err := calendarRepo.FindEvent(c.Request().Context(), cEvent)
		if err != nil {
			return err
		}

		return Render(c, http.StatusOK, components.CalendarEventActions(event))
	}
}
