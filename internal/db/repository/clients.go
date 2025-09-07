package repository

import (
	"context"

	"github.com/Gromp-13/IronGym/internal/models"
)

// Получение списка всех клиентов
func (repo *PGRepo) GetClients() ([]models.Client, error) {
	rows, err := repo.pool.Query(context.Background(),
		`SELECT id, lastname, firstname, middlename, phonenumber, birthdate, cardbarcode, gender, status_id 
   FROM clients;`)
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
			&item.Gender,
			&item.StatusID,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

// Добавление нового клиента
func (repo *PGRepo) NewClient(item models.Client) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, err := repo.pool.Exec(context.Background(),
		`INSERT INTO clients (lastname, firstname, middlename, phonenumber, birthdate, cardbarcode, gender, status_id) 
   VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		item.LastName,
		item.FirstName,
		item.MiddleName,
		item.PhoneNumber,
		item.BirthDate,
		item.CardBarcode,
		item.Gender,
		item.StatusID,
	)
	return err
}

// Получение клиента по ID
func (repo *PGRepo) GetClientByID(id int32) (models.Client, error) {
	var c models.Client
	err := repo.pool.QueryRow(context.Background(),
		`SELECT id, lastname, firstname, middlename, phonenumber, birthdate, cardbarcode, gender, status_id 
   FROM clients WHERE id = $1`,
		id,
	).Scan(
		&c.ID,
		&c.LastName,
		&c.FirstName,
		&c.MiddleName,
		&c.PhoneNumber,
		&c.BirthDate,
		&c.CardBarcode,
		&c.Gender,
		&c.StatusID,
	)

	return c, err
}

// Закрытие подключения
func (repo *PGRepo) Close() {
	repo.pool.Close()
}

func (repo *PGRepo) GetLastClientByBarcode(barcode string) (models.Client, error) {
	var c models.Client
	err := repo.pool.QueryRow(context.Background(),
		`SELECT id, lastname, firstname, middlename, phonenumber, birthdate, cardbarcode, gender, statusid
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
		&c.StatusID,
	)
	return c, err
}
