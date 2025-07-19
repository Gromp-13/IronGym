package repository

import (
	"context"

	"github.com/Gromp-13/IronGym/internal/models"
)

func (repo *PGRepo) GetClients() ([]models.Clients, error) {
	rows, err := repo.pool.Query(context.Background(), `
	SELECT id, lastname, firstname, middlename, phonenumber, birthdate, cardbarcode 
	FROM clients;
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.Clients
	for rows.Next() {
		var item models.Clients
		err = rows.Scan(
			&item.ID,
			&item.LastName,
			&item.FirstName,
			&item.MiddleName,
			&item.PhoneNumber,
			&item.BirthDate,
			&item.CardBarcode,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, item)
	}

	return data, nil
}
