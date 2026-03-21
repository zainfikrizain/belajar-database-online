package repositories

import "database/sql"

type ProductRepository struct {
	db *sql.DB
}

func newProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
