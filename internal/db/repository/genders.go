package repository

import (
	"context"

	"github.com/Gromp-13/IronGym/internal/models"
)

func (repo *PGRepo) GetGenders() ([]models.Gender, error) {
	rows, err := repo.pool.Query(context.Background(), `
	SELECT id, description FROM genders ORDER BY id;
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genders []models.Gender
	for rows.Next() {
		var g models.Gender
		err := rows.Scan(&g.ID, &g.Description)
		if err != nil {
			return nil, err
		}
		genders = append(genders, g)
	}
	return genders, nil
}
