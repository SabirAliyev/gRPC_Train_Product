package postgresql

import (
	"context"
	"example.com/go-productmgmt-grpc/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductModel struct {
	Pool *pgxpool.Pool
}

func (m *ProductModel) FindById(id int32) (models.Product, error) {
	query := `SELECT Id, name, description, value FROM "product" WHERE Id = $1`

	var prod models.Product

	if err := m.Pool.QueryRow(context.Background(), query, id).
		Scan(&prod.Id, &prod.Name, &prod.Description, &prod.Value); err != nil {
		return models.Product{}, err
	}

	return prod, nil
}

func (m *ProductModel) Insert(prod *models.Product) (int32, error) {
	query := `INSERT INTO product (id, name, description, value) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int32

	if err := m.Pool.QueryRow(context.Background(), query, prod.Name, prod.Description, prod.Value).
		Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
