package domain

import (
	"fmt"
	"strings"
	"time"
)

type (
	CalendarType string

	Calendar struct {
		ID           string `yaml:"id"`
		Name         string `yaml:"name"` // the name is for representing in the ui
		Color        string `yaml:"color"`
		Icon         string `yaml:"icon"`
		Type         CalendarType
		RenderOption CalendarRenderOption
	}

	CalendarRenderOption struct {
		StartHour int
		EndHour   int
	}

	CalendarYear struct{}

	CalendarMonth struct{}

	CalendarWeek struct {
		Dates  []time.Time
		Events []CalendarEventContract
		Days   []CalendarDay
	}

	CalendarDay struct {
		Events []CalendarEventContract
	}

	CalendarEventContract interface {
		ID() CalendarEventID
		Name() string
		Description() string
		Start() time.Time
		End() time.Time
		SetRegistrations([]Registration)
		Registrations() []Registration
	}

	CalendarEventID struct {
		RefType string
		RefID   string
	}
)

func NewCalendarEventID(refType string, refID string) CalendarEventID {
	return CalendarEventID{
		RefType: refID,
		RefID:   refID,
	}
}

func (c CalendarEventID) ToString() string {
	return fmt.Sprintf("%s-%s-%s", c.RefType, c.RefID, "something")
}

func (c *CalendarEventID) FromString(input string) {
	fmt.Println(input)
	parts := strings.Split(input, "-")
	fmt.Println(parts)
	c.RefType = parts[0]
	c.RefID = parts[1]
}
