package infrastructure

import (
	"context"
	"errors"
	"time"

	"github.com/zenginechris/studio/domain"
)

type CalendarRepository struct {
	classRepo        ClassRepository
	registrationRepo RegistrationRepository
}

func NewCalendarRepository(
	classRepo ClassRepository,
	registrationRepo RegistrationRepository,
) CalendarRepository {
	return CalendarRepository{
		classRepo,
		registrationRepo,
	}
}

func (r CalendarRepository) FindEvent(ctx context.Context, id domain.CalendarEventID) (domain.CalendarEventContract, error) {
	class, err := r.classRepo.Find(ctx, id.RefID)
	if err != nil {
		return nil, err
	}

	events := class.ToCalendarEvents([]time.Time{
		time.Now(),
	})

	if len(events) < 1 {
		return nil, errors.New("no events")
	}

	// search for the registrations
	registrations, err := r.registrationRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	result := events[0]
	result.SetRegistrations(registrations)

	return result, nil
}

func (r CalendarRepository) GetEvents(ctx context.Context) ([]domain.CalendarEventContract, error) {
	classes, err := r.classRepo.Get()
	if err != nil {
		return nil, err
	}
	result := []domain.CalendarEventContract{}

	for _, c := range classes {
		result = append(result, c.ToCalendarEvents([]time.Time{time.Now()})...)
	}
	return result, nil
}
