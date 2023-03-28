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
		balance INT,
		created_at timestamp

	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateAccount(acc *Account) error {

	query := `insert into account (first_name,last_name,number,balance,created_at) VALUES ($1,$2,$3,$4,$5)`

	resp, err := s.db.Query(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)
	return nil
}

func (db *PostgressStore) DeleteAccount(id int) error {

	return nil
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

func tableScan(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt)

	return account, err
}
