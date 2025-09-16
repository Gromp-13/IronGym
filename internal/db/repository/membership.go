package repository

import (
	"context"

	"github.com/Gromp-13/IronGym/internal/models"
)

// Добавление нового типа абонемента в справочник memberships
func (repo *PGRepo) NewMembership(item models.Membership) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, err := repo.pool.Exec(
		context.Background(),
		`INSERT INTO memberships (name, price, duration_days)
   VALUES ($1, $2, $3)`,
		item.Name,
		item.Price,        // int32
		item.DurationDays, // int32
	)
	return err
}

// Получение списка всех типов абонементов
func (repo *PGRepo) GetMemberships() ([]models.Membership, error) {
	rows, err := repo.pool.Query(
		context.Background(),
		`SELECT id, name, price::int, duration_days
     FROM memberships
  ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.Membership
	for rows.Next() {
		var item models.Membership
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Price,        // int32
			&item.DurationDays, // int32
		); err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

// Обновление цены абонемента
func (repo *PGRepo) UpdateMembershipPrice(id int32, newPrice int32) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, err := repo.pool.Exec(
		context.Background(),
		`UPDATE memberships
      SET price = $1
    WHERE id = $2`,
		newPrice,
		id,
	)

	return err
}

func (repo *PGRepo) GetMembershipPrice(id int32) (int32, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var price int32
	err := repo.pool.QueryRow(
		context.Background(),
		`SELECT price FROM memberships WHERE id = $1`,
		id,
	).Scan(&price)
	if err != nil {
		return 0, err
	}

	return price, nil
}
