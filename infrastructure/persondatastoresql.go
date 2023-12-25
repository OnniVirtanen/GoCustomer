package infrastructure

import (
	"database/sql"

	"example.com/backend/core/model/aggregate"
	"example.com/backend/core/model/entity"
	"github.com/google/uuid"
)

// CustomerRepository holds the database connection.
type CustomerRepository struct {
	DB *sql.DB
}

// NewCustomerRepository creates a new instance of CustomerRepository.
func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

// Get retrieves a single customer by their UUID.
func (cr *CustomerRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	var c aggregate.Customer
	c.Person = &entity.Person{}

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			phone,
			country_code,
			street,
			postal_code,
			city,
			age
		FROM person WHERE id = ?
	`

	err := cr.DB.QueryRow(query, id).Scan(
		&c.Person.ID, &c.Person.FirstName, &c.Person.LastName, &c.Person.Email, &c.Person.Phone,
		&c.Person.CountryCode, &c.Person.Street, &c.Person.PostalCode, &c.Person.City, &c.Person.Age,
	)

	if err != nil {
		return aggregate.Customer{}, err
	}
	return c, nil
}

// GetAll retrieves all customers from the database.
func (cr *CustomerRepository) GetAll() ([]aggregate.Customer, error) {
	var customers []aggregate.Customer

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			phone,
			country_code,
			street,
			postal_code,
			city,
			age
		FROM person
	`

	rows, err := cr.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c aggregate.Customer
		c.Person = &entity.Person{}

		err := rows.Scan(
			&c.Person.ID, &c.Person.FirstName, &c.Person.LastName, &c.Person.Email, &c.Person.Phone,
			&c.Person.CountryCode, &c.Person.Street, &c.Person.PostalCode, &c.Person.City, &c.Person.Age,
		)

		if err != nil {
			return nil, err
		}

		customers = append(customers, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

// Add inserts a new customer into the database.
func (cr *CustomerRepository) Add(customer aggregate.Customer) error {
	query := `
		INSERT INTO person (
			id, first_name, last_name, email, phone, country_code, street, postal_code, city, age
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := cr.DB.Exec(query,
		customer.Person.ID, customer.Person.FirstName, customer.Person.LastName, customer.Person.Email,
		customer.Person.Phone, customer.Person.CountryCode, customer.Person.Street, customer.Person.PostalCode,
		customer.Person.City, customer.Person.Age,
	)

	return err
}

// Update modifies an existing customer in the database.
func (cr *CustomerRepository) Update(customer aggregate.Customer, id uuid.UUID) error {
	query := `
		UPDATE person SET
			first_name = ?, last_name = ?, email = ?, phone = ?, 
			country_code = ?, street = ?, postal_code = ?, city = ?, age = ?
		WHERE id = ?
	`

	_, err := cr.DB.Exec(query,
		customer.Person.FirstName, customer.Person.LastName, customer.Person.Email, customer.Person.Phone,
		customer.Person.CountryCode, customer.Person.Street, customer.Person.PostalCode, customer.Person.City,
		customer.Person.Age, id,
	)

	return err
}

// Delete removes a customer from the database.
func (cr *CustomerRepository) Delete(id uuid.UUID) error {
	query := "DELETE FROM person WHERE id = ?"
	_, err := cr.DB.Exec(query, id)
	return err
}
