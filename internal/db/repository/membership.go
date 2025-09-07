package repository

import (
	"context"
	"time"

	"github.com/Gromp-13/IronGym/internal/models"
)

// Добавление нового типа абонемента в справочник memberships
func (repo *PGRepo) NewMembership(item models.Membership) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, err := repo.pool.Exec(context.Background(),
		`INSERT INTO memberships (name, price, duration_days)
   VALUES ($1, $2, $3)`,
		item.Name,
		item.Price,        // int32 -> numeric в БД сконвертится неявно
		item.DurationDays, // int32
	)
	return err
}

// Получение списка всех типов абонементов
func (repo *PGRepo) GetMemberships() ([]models.Membership, error) {
	rows, err := repo.pool.Query(context.Background(),
		`SELECT id, name, price::int, duration_days
     FROM memberships
     ORDER BY id`)
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

// Привязка абонемента к клиенту (покупка/назначение)
func (repo *PGRepo) AssignMembershipToClient(
	clientID, membershipID, paidAmount, statusID int32,
	startDate time.Time,
) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, err := repo.pool.Exec(context.Background(),
		`INSERT INTO client_memberships
   (client_id, membership_id, start_date, paid_amount, created_at, status_id)
   VALUES ($1,        $2,           $3,         $4,          $5,        $6)`,
		clientID,
		membershipID,
		startDate,  // date/timestamp совместим (Postgres приведёт к date при необходимости)
		paidAmount, // int32 -> numeric
		time.Now(),
		statusID,
	)
	return err
}

// Получение всех абонементов конкретного клиента
func (repo *PGRepo) GetClientMemberships(clientID int32) ([]models.ClientMembership, error) {
	rows, err := repo.pool.Query(context.Background(),
		`SELECT id,
          client_id,
          membership_id,
          start_date,
          paid_amount::int,
          created_at,
          status_id
     FROM client_memberships
    WHERE client_id = $1
    ORDER BY id`,
		clientID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.ClientMembership
	for rows.Next() {
		var item models.ClientMembership
		if err := rows.Scan(
			&item.ID,
			&item.ClientID,
			&item.MembershipID,
			&item.StartDate,
			&item.PaidAmount, // int32
			&item.CreatedAT,
			&item.StatusID,
		); err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

func (repo *PGRepo) NewClientMembership(cm models.ClientMembership) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, err := repo.pool.Exec(context.Background(),
		`INSERT INTO client_memberships (clientid, membershipid, startdate, paidamount, createdat, statusid)
        VALUES ($1, $2, $3, $4, $5, $6)`,
		cm.ClientID,
		cm.MembershipID,
		cm.StartDate,
		cm.PaidAmount,
		cm.CreatedAT,
		cm.StatusID,
	)

	return err
}
