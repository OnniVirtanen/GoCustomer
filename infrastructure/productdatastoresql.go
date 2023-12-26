package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	"example.com/backend/core/model/entity"
	"example.com/backend/core/repository"
	"github.com/google/uuid"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Get(id uuid.UUID) (entity.Product, error) {
	var p entity.Product

	query := `
		SELECT
			name,
			description,
			quantity,
			price,
			image,
			category,
			createdAt,
			updatedAt,
			discount
		FROM product WHERE id = ?
	`

	err := r.DB.QueryRow(query, id).Scan(
		&p.ID, &p.Name, &p.Description, &p.Quantity, &p.Price, &p.Image,
		&p.Category, &p.CreatedAt, &p.UpdatedAt, &p.Discount,
	)

	if err != nil {
		return entity.Product{}, err
	}
	return p, nil
}

func (r *ProductRepository) GetAll() ([]entity.Product, error) {
	fmt.Println("-2")
	var products []entity.Product

	query := `
		SELECT
			id,
			name,
			description,
			quantity,
			price,
			image,
			category,
			createdAt,
			updatedAt,
			discount
		FROM product
	`
	fmt.Println("-1")
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fmt.Println("0")

	for rows.Next() {
		fmt.Println("1")
		var p entity.Product

		err := rows.Scan(
			&p.ID, &p.Name, &p.Description, &p.Quantity, &p.Price, &p.Image,
			&p.Category, &p.CreatedAt, &p.UpdatedAt, &p.Discount,
		)
		fmt.Println("err", err)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	fmt.Println("2")

	if err = rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("3")

	fmt.Println("products at productdatastoresql.go GetAll()", products)

	return products, nil
}

func (r *ProductRepository) Add(product entity.Product) error {
	query := `
		INSERT INTO product (
			id,
			name,
			description,
			quantity,
			price,
			image,
			category,
			createdAt,
			discount
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	fmt.Println("In productdatastoresql.go.add")

	_, err := r.DB.Exec(query,
		product.ID, product.Name, product.Description, product.Quantity, product.Price,
		product.Image, product.Category, time.Now(), product.Discount,
	)

	return err
}

func (r *ProductRepository) Update(product entity.Product, id uuid.UUID) error {
	query := `
		UPDATE product SET
			name = ?,
			description = ?,
			quantity = ?,
			price = ?,
			image = ?,
			category = ?,
			updatedAt = ?,
			discount = ?
		WHERE id = ?
	`

	_, err := r.DB.Exec(query,
		product.Name, product.Description, product.Quantity,
		product.Price, product.Image, product.Category, time.Now(), product.Discount, id,
	)

	return err
}

func (r *ProductRepository) Delete(id uuid.UUID) error {
	var exists int
	checkQuery := "SELECT 1 FROM product WHERE id = ?"
	err := r.DB.QueryRow(checkQuery, id).Scan(&exists)

	if err == sql.ErrNoRows {
		return fmt.Errorf("product does not exist: %w", repository.ErrProductNotFound)
	} else if err != nil {
		// Some other error occurred
		return err
	}
	deleteQuery := "DELETE FROM product WHERE id = ?"
	_, err = r.DB.Exec(deleteQuery, id)
	return err
}
