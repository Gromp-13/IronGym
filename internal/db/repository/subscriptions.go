package repository

import (
	"context"

	"github.com/Gromp-13/IronGym/internal/models"
)

func (repo *PGRepo) NewSubscription(sub models.Subscriptions) error {
	_, err := repo.pool.Exec(context.Background(), `
	INSERT INTO subscriptions (clientid, startdate, durationdays, enddate)
	VALUES ($1, $2, $3, $4)`,
		sub.ClientID,
		sub.StartDate,
		sub.DurationDays,
		sub.EndDate,
	)
	return err
}

func (repo *PGRepo) GetLastClientByBarcode(barcode string) (models.Client, error) {
	var c models.Client
	err := repo.pool.QueryRow(context.Background(), `
	SELECT id, lastname, firstname, middlename, phonenumber, birthdate, cardbarcode, gender
	FROM clients
	WHERE cardbarcode = $1
	ORDER BY id DESC
	LIMIT 1
	`, barcode).Scan(
		&c.ID,
		&c.LastName,
		&c.FirstName,
		&c.MiddleName,
		&c.PhoneNumber,
		&c.BirthDate,
		&c.CardBarcode,
		&c.Gender,
	)
	return c, err
}
