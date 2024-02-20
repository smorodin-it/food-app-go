package services

import (
	"food-backend/src/domains"
	"food-backend/src/forms"
	"food-backend/src/repositories"
	"github.com/google/uuid"
	"time"
)

type MeasurementService interface {
	ListByUserId(page int, perPage int, userId string) (measurements []domains.Measurement, err error)
	Create(mf forms.MeasurementCreateForm) (id *string, err error)
	Retrieve(measurementId string) (measurement *domains.Measurement, err error)
	Update(mf forms.MeasurementUpdateForm, measurementId string) (err error)
	Delete(measurementId string) (err error)
}

type measurementService struct {
	r repositories.MeasurementRepository
}

func (s measurementService) ListByUserId(page int, perPage int, userId string) (measurements []domains.Measurement, err error) {
	measurements, err = s.r.ListByUserId(page, perPage, userId)
	if err != nil {
		return nil, err
	}

	return measurements, nil
}

func (s measurementService) Create(mf forms.MeasurementCreateForm) (id *string, err error) {
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

func (s measurementService) Retrieve(measurementId string) (measurement *domains.Measurement, err error) {
	measurement, err = s.r.Retrieve(measurementId)
	if err != nil {
		return nil, err
	}

	return measurement, nil
}

func (s measurementService) Update(mf forms.MeasurementUpdateForm, measurementId string) (err error) {
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

func (s measurementService) Delete(measurementId string) (err error) {
	err = s.r.Delete(measurementId)
	if err != nil {
		return err
	}

	return nil
}

func NewMeasurementService(repository repositories.MeasurementRepository) MeasurementService {
	return &measurementService{r: repository}
}
