package infrastructure

import (
	"context"
	"errors"

	"github.com/zenginechris/studio/domain"
)

type RegistrationRepository struct {
	items []domain.Registration
}

func NewRegistrationRepository() RegistrationRepository {
	return RegistrationRepository{
		items: make([]domain.Registration, 0),
	}
}

func (r RegistrationRepository) Find(ctx context.Context, id domain.CalendarEventID) (domain.Registration, error) {
	for _, item := range r.items {
		if item.ReferenceID.ToString() == id.ToString() {
			return item, nil
		}
	}
	return domain.Registration{}, errors.New("registration not found")
}

func (r RegistrationRepository) Get(ctx context.Context, id domain.CalendarEventID) ([]domain.Registration, error) {
	result := []domain.Registration{}
	for _, item := range r.items {
		if item.ReferenceID.ToString() == id.ToString() {
			result = append(result, item)
		}
	}
	return result, nil
}

func (r *RegistrationRepository) Store(ctx context.Context, registration domain.Registration) (domain.Registration, error) {
	if registration.ID == "" {
		registration.ID = "outou"
	}
	r.items = append(r.items, registration)
	return registration, nil
}
