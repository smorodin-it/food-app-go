package services

import (
	"food-backend/src/domains"
	"food-backend/src/forms"
	"food-backend/src/repositories"
	"github.com/google/uuid"
	"time"
)

type MeasurementService struct {
	r repositories.MeasurementRepository
}

func (s MeasurementService) ListByUserId(page int, perPage int, userId string) (measurements []domains.Measurement, err error) {
	measurements, err = s.r.ListByUserId(page, perPage, userId)
	if err != nil {
		return nil, err
	}

	return measurements, nil
}

func (s MeasurementService) Create(mf forms.MeasurementCreateForm) (id *string, err error) {
	m := &domains.Measurement{
		MeasurementId:     uuid.New().String(),
		UserId:            mf.UserId,
		MeasurementWeight: mf.MeasurementWeight,
		Date:              time.Now(),
	}
	err = s.r.Create(m)
	if err != nil {
		return nil, err
	}

	return &m.MeasurementId, nil
}

func (s MeasurementService) Update(mf forms.MeasurementUpdateForm, measurementId string) (err error) {
	m := &domains.Measurement{
		MeasurementId:     measurementId,
		MeasurementWeight: mf.MeasurementWeight,
	}
	err = s.r.Update(m)
	if err != nil {
		return err
	}

	return nil
}

func (s MeasurementService) Delete(measurementId string) (err error) {
	err = s.r.Delete(measurementId)
	if err != nil {
		return err
	}

	return nil
}
