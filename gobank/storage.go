package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	GetAccount(id int) (*Account, error)
	CreateAccount(account *CreateAccountRequest) error
	DeleteAccount(id int) error
	UpdateAccpunt(account *Account) error
	Transfer(fromID, toID int, amount float64) error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(dbUrl string) (*PostgresStorage, error) {
	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		return nil, err
	}

	return &PostgresStorage{db: conn}, nil
}

func (s *PostgresStorage) Init() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		balance FLOAT,
		acc_number TEXT NOT NULL
	)`)

	return err
}

func (s *PostgresStorage) Close() {
	s.db.Close()
}

func (s *PostgresStorage) GetAccount(id int) (*Account, error) {
	account, err := s.db.Query(`SELECT id, name, balance, acc_number FROM accounts WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	accountData := &Account{}
	if account.Next() {
		err = account.Scan(&accountData.ID, &accountData.Fullname, &accountData.Balance, &accountData.AccNumber)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, sql.ErrNoRows
	}

	return accountData, nil
}

func (s *PostgresStorage) CreateAccount(account *CreateAccountRequest) error {
	_, err := s.db.Exec(
		`INSERT INTO accounts 
		(name, balance, acc_number) 
		VALUES ($1, 0, encode(sha256 (random()::text::bytea), 'hex'))	 
		`,
		account.Fullname,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStorage) UpdateAccpunt(account *Account) error {
	return nil
}

func (s *PostgresStorage) Transfer(fromID, toID int, amount float64) error {
	return nil
}
