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
func (m *DBModel) ListResources() ([]*Resource, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `select id, resource_name, resource_id, location from resources order by index`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*Resource

	for rows.Next() {
		var r Resource
		err := rows.Scan(&r.ID, &r.ResourceName, &r.ResourceID, &r.Location)
		if err != nil {
			return nil, err
		}
		res = append(res, &r)
	}

	return res, nil
}
func (m *DBModel) PostMovie()   {}
func (m *DBModel) DeleteMovie() {}
