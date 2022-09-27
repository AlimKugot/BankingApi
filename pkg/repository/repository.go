package main

import "github.com/jmoiron/sqlx"

type UserAccount interface {
}

type Repository struct {
	UserAccount
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
