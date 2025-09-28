package domain

import "time"

type (
	Interval struct {
		Weekdays []time.Weekday `yaml:"weekdays"`
		Start    time.Time      `yaml:"start"`
		End      time.Time      `yaml:"end"`
	}

	Class struct {
		ID string
		// Calendar    Calendar
		Name        string
		Description string

		Interval  Interval
		StartDate time.Time
		EndDate   time.Time

		CreatedAt time.Time
		UpdatedAt time.Time
	}

	ClassAppointment struct {
		registrations []Registration
		class         Class
		start         time.Time
		end           time.Time
	}
)

func (a ClassAppointment) Name() string {
	return a.class.Name
}

func (a ClassAppointment) Start() time.Time {
	return a.start
}

func (a ClassAppointment) End() time.Time {
	return a.end
}

func (a ClassAppointment) Description() string {
	return ""
}

func (a ClassAppointment) ID() CalendarEventID {
	return NewCalendarEventID("cl", a.class.ID)
}

func (a *ClassAppointment) SetRegistrations(registrations []Registration) {
	a.registrations = append(a.registrations, registrations...)
}

func (a ClassAppointment) Registrations() []Registration {
	return a.registrations
}

func (c Class) ToCalendarEvents(dates []time.Time) []CalendarEventContract {
	var cEvents []CalendarEventContract
	for _, date := range dates {
		for _, day := range c.Interval.Weekdays {
			// now we can check and generate the missing data
			if day == date.Weekday() {
				cEvents = append(cEvents, &ClassAppointment{
					class: c,
					start: time.Date(
						date.Year(),
						date.Month(),
						date.Day(),
						c.Interval.Start.Hour(),
						c.Interval.Start.Minute(),
						0,
						0,
						time.Local,
					),
					end: time.Date(
						date.Year(),
						date.Month(),
						date.Day(),
						c.Interval.End.Hour(),
						c.Interval.End.Minute(),
						0,
						0,
						time.Local,
					),
				})
			}
		}
	}
	return cEvents
}

// we want to convert a class into calendar events
