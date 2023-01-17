package db

import (
	customer "customers"
	"database/sql"
	"fmt"
)

type CustomerStore struct {
	*sql.DB
}

func (s *CustomerStore) Customer(id int) (customer.Customer, error) {
	defer s.Close()
	query := `select user_id, email, password, created_on, last_login from customers where user_id = $1`
	var c customer.Customer
	row := s.QueryRow(query, id)
	if err := row.Scan(&c.ID, &c.Email, &c.First_Name, &c.Last_Name); err != nil {
		return customer.Customer{}, fmt.Errorf("error getting customer: %w", err)
	}
	return c, nil
}

func (s *CustomerStore) Customers() ([]customer.Customer, error) {
	defer s.Close()
	query := `select * from customers`
	var cc []customer.Customer
	rows, err := s.Query(query)
	if err != nil {
		return []customer.Customer{}, fmt.Errorf("error getting customer: %w", err)
	}

	for rows.Next() {
		var c customer.Customer
		if err := rows.Scan(&cc); err != nil {
			return []customer.Customer{}, fmt.Errorf("error getting customer: %w", err)
		}
		cc = append(cc, c)
	}
	return cc, nil
}

func (s *CustomerStore) CreateCustomer(c *customer.Customer) error {
	defer s.Close()
	query := `insert into customers (user_id, first_name, last_name, email) values ($1, $2, $3, $4)`
	if _, err := s.Exec(query, &c.ID, &c.First_Name, &c.Last_Name, &c.Email); err != nil {
		return err
	}
	return nil
}

func (s *CustomerStore) UpdateCustomer(c *customer.Customer) error {
	defer s.Close()
	query := `update customers set first_name = $1, last_name = $2, email=$3 where user_id = $4`
	if _, err := s.Exec(query, &c.First_Name, &c.Last_Name, &c.Email, &c.ID); err != nil {
		return err
	}
	return nil
}

func (s *CustomerStore) DeleteCustomer(c *customer.Customer) error {
	defer s.Close()
	query := `delete from customers where user_id = $1`
	if _, err := s.Exec(query, &c.ID); err != nil {
		return err
	}
	return nil
}
