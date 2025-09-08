package repository

import (
	"context"
	"time"

	"github.com/Gromp-13/IronGym/internal/models"
)

// Привязка абонемента к клиенту (покупка/назначение)
func (repo *PGRepo) AssignMembershipToClient(
	clientID, membershipID, paidAmount, statusID int32,
	startDate time.Time,
) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, err := repo.pool.Exec(
		context.Background(),
		`INSERT INTO client_memberships
   (client_id, membership_id, start_date, paid_amount, created_at, status_id)
   VALUES ($1, $2, $3, $4, $5, $6)`,
		clientID,
		membershipID,
		startDate,
		paidAmount,
		time.Now(),
		statusID,
	)
	return err
}

// Получение всех абонементов конкретного клиента
func (repo *PGRepo) GetClientMemberships(clientID int32) ([]models.ClientMembership, error) {
	rows, err := repo.pool.Query(
		context.Background(),
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
			&item.PaidAmount,
			&item.CreatedAT,
			&item.StatusID,
		); err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

// Добавление записи о новом абонементе клиента
func (repo *PGRepo) NewClientMembership(cm models.ClientMembership) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, err := repo.pool.Exec(
		context.Background(),
		`INSERT INTO client_memberships
   (client_id, membership_id, start_date, paid_amount, created_at, status_id)
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
