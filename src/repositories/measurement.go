package repositories

import (
	"food-backend/src/database"
	"food-backend/src/domains"
	"food-backend/src/utils"
)

type MeasurementRepository struct {
}

func (r MeasurementRepository) ListByUserId(page int, perPage int, userId string) (measurements []domains.Measurement, err error) {
	limit, offset, withPagination := utils.CalcPagination(page, perPage)

	if withPagination {
		sql := "SELECT * FROM measurements WHERE measurements.user_id = $1 ORDER BY date DESC  LIMIT $2 OFFSET $3"
		err = database.DBCon.Select(&measurements, sql, userId, limit, offset)
	} else {
		sql := "SELECT * FROM measurements WHERE measurements.user_id = $1 ORDER BY date DESC"
		err = database.DBCon.Select(&measurements, sql, userId)
	}

	if err != nil {
		return nil, err
	}

	return measurements, nil
}

func (r MeasurementRepository) Create(measurement *domains.Measurement) (err error) {
	sql := "INSERT INTO measurements (measurement_id, user_id, measurement_weight, date) VALUES ($1, $2, $3, $4)"
	_, err = database.DBCon.Exec(sql, measurement.MeasurementId, measurement.UserId, measurement.MeasurementWeight, measurement.Date)
	if err != nil {
		return err
	}

	return nil
}

func (r MeasurementRepository) Update(measurement *domains.Measurement) error {
	sql := "UPDATE measurements SET measurement_weight = $1 WHERE measurement_id = $2"
	_, err := database.DBCon.Exec(sql, measurement.MeasurementWeight, measurement.MeasurementId)
	if err != nil {
		return err
	}

	return nil
}

func (r MeasurementRepository) Delete(measurementId string) error {
	sql := "DELETE FROM measurements where measurement_id = $1"
	_, err := database.DBCon.Exec(sql, measurementId)
	if err != nil {
		return err
	}

	return nil
}
