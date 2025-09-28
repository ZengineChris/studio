package infrastructure

import (
	"context"
	"time"

	"github.com/zenginechris/studio/domain"
)

type ClassRepository struct{}

func NewClassRepository() ClassRepository {
	return ClassRepository{}
}

func (r ClassRepository) Get() ([]domain.Class, error) {
	result := []domain.Class{
		{
			ID: "tee",
			Interval: domain.Interval{
				Weekdays: []time.Weekday{time.Monday, time.Saturday},
				Start:    time.Now(),
				End:      time.Now().Add(2 * time.Hour),
			},
			Name:      "bjj",
			StartDate: time.Now().Add(-100 * time.Hour),
			EndDate:   time.Now().Add(100 * time.Hour),
		},
	}
	return result, nil
}

func (r ClassRepository) Find(ctx context.Context, id string) (domain.Class, error) {
	result := domain.Class{
		ID: "tee",
		Interval: domain.Interval{
			Weekdays: []time.Weekday{time.Monday, time.Saturday},
			Start:    time.Now(),
			End:      time.Now().Add(2 * time.Hour),
		},
		Name:      "bjj",
		StartDate: time.Now().Add(-100 * time.Hour),
		EndDate:   time.Now().Add(100 * time.Hour),
	}

	return result, nil
}
