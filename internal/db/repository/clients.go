package repository

import (
	"context"

	"github.com/Gromp-13/IronGym/internal/models"
)

func (repo *PGRepo) GetClient() ([]models.Client, error) {
	rows, err := repo.pool.Query(context.Background(), `
	SELECT id, lastname, firstname, middlename, phonenumber, birthdate, cardbarcode, gender 
	FROM clients;
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.Client
	for rows.Next() {
		var item models.Client
		err = rows.Scan(
			&item.ID,
			&item.LastName,
			&item.FirstName,
			&item.MiddleName,
			&item.PhoneNumber,
			&item.BirthDate,
			&item.CardBarcode,
			&item.GenderID,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, item)
	}

	return data, nil
}

func (repo *PGRepo) NewClients(item models.Client) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	_, err := repo.pool.Exec(context.Background(), `
	INSERT INTO clients (lastname, firstname, middlename, phonenumber, birthdate, cardbarcode)
	VALUES ($1, $2, $3, $4, $5, $6)`,
		item.LastName,
		item.FirstName,
		item.MiddleName,
		item.PhoneNumber,
		item.BirthDate,
		item.CardBarcode,
	)
	return err
}

func (repo *PGRepo) Close() {
	repo.pool.Close()
}
