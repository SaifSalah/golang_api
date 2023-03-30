package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type IStorage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
	GetAccountByNumber(int) (*Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func PostgresConn() (*PostgressStore, error) {

	connStr := "postgres://postgres:123123@localhost/uat_testing?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil
}

func (s *PostgressStore) Init() error {

	return s.CreateAccountTable()
}

func (s *PostgressStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (

		id serial PRIMARY KEY,
		first_name varchar(50),
		last_name varchar(50),
		number INT,
		password varchar(255),
		balance INT,
		created_at timestamp

	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateAccount(acc *Account) error {

	query := `insert into account (first_name,last_name,number,balance,password,created_at) VALUES ($1,$2,$3,$4,$5,$6)`

	_, err := s.db.Query(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.Password, acc.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (db *PostgressStore) DeleteAccount(id int) error {

	_, err := db.db.Query("DELETE FROM account WHERE id = $1", id)
	return err

}

func (db *PostgressStore) UpdateAccount(*Account) error {

	return nil
}

func (db *PostgressStore) GetAccountByID(id int) (*Account, error) {

	rows, err := db.db.Query("SELECT * FROM account WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return tableScan(rows)
	}
	return nil, fmt.Errorf("Account %d not found", id)
}

func (db *PostgressStore) GetAccounts() ([]*Account, error) {

	rows, err := db.db.Query("SELECT * FROM account")

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account, err := tableScan(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil

}

func (db *PostgressStore) GetAccountByNumber(number int) (*Account, error) {

	rows, err := db.db.Query("SELECT * FROM account WHERE number = $1", number)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return tableScan(rows)
	}
	return nil, fmt.Errorf("Account %d not found", number)

}
