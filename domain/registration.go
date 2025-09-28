package domain

type (
	RegistrationState string

	Registration struct {
		ID          string
		ReferenceID CalendarEventID
		UserID      string
		State       RegistrationState
	}
)

// this is in the end handled by a state machine
const (
	// the registration is created and is waiting for all hooks to be true
	RegistrationStateCreated RegistrationState = "pending"

	// the registration is complete and the spot is booked
	RegistrationStateConfirmed RegistrationState = "confirmed"

	// the user has checked into the class
	RegistrationStateCheckedIn RegistrationState = "checked_in"
)
