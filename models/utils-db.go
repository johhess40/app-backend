package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) GetResource(id int) (*Resource, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `select id, resource_name, resource_id, location from resources where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var res Resource

	err := row.Scan(&res.ID, &res.ResourceName, &res.ResourceID, &res.Location)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
func (m *DBModel) ListMovies() ([]*Movie, error) {
	return nil, nil
}
func (m *DBModel) PostMovie()   {}
func (m *DBModel) DeleteMovie() {}
