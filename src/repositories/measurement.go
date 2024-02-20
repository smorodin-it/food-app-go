package repositories

import (
	"food-backend/src/domains"
	"food-backend/src/utils"
	"github.com/jmoiron/sqlx"
)

type MeasurementRepository interface {
	ListByUserId(page int, perPage int, userId string) (measurements []domains.Measurement, err error)
	Create(measurement *domains.Measurement) (err error)
	Retrieve(measurementId string) (measurement *domains.Measurement, err error)
	Update(measurement *domains.Measurement) (err error)
	Delete(measurementId string) (err error)
}

type measurementRepository struct {
	db *sqlx.DB
}

func (r measurementRepository) ListByUserId(page int, perPage int, userId string) (measurements []domains.Measurement, err error) {
	limit, offset, withPagination := utils.CalcPagination(page, perPage)

	if withPagination {
		sql := "SELECT * FROM measurements WHERE measurements.user_id = $1 ORDER BY date DESC  LIMIT $2 OFFSET $3"
		err = r.db.Select(&measurements, sql, userId, limit, offset)
	} else {
		sql := "SELECT * FROM measurements WHERE measurements.user_id = $1 ORDER BY date DESC"
		err = r.db.Select(&measurements, sql, userId)
	}

	if err != nil {
		return nil, err
	}

	return measurements, nil
}

func (r measurementRepository) Create(measurement *domains.Measurement) (err error) {
	sql := "INSERT INTO measurements (measurement_id, user_id, measurement_weight, date) VALUES ($1, $2, $3, $4)"
	_, err = r.db.Exec(sql, measurement.MeasurementId, measurement.UserId, measurement.MeasurementWeight, measurement.Date)
	if err != nil {
		return err
	}

	return nil
}

func (r measurementRepository) Retrieve(measurementId string) (measurement *domains.Measurement, err error) {
	sql := "SELECT * from measurements WHERE measurement_id = $1"
	err = r.db.Get(measurement, sql, measurementId)
	if err != nil {
		return nil, err
	}

	return measurement, nil
}

func (r measurementRepository) Update(measurement *domains.Measurement) (err error) {
	sql := "UPDATE measurements SET measurement_weight = $1 WHERE measurement_id = $2"
	_, err = r.db.Exec(sql, measurement.MeasurementWeight, measurement.MeasurementId)
	if err != nil {
		return err
	}

	return nil
}

func (r measurementRepository) Delete(measurementId string) (err error) {
	sql := "DELETE FROM measurements where measurement_id = $1"
	_, err = r.db.Exec(sql, measurementId)
	if err != nil {
		return err
	}

	return nil
}

func NewMeasurementRepository(db *sqlx.DB) MeasurementRepository {
	return &measurementRepository{db: db}
}
