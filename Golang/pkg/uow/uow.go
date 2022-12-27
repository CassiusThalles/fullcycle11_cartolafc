package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow *Uow) error) error
	CommitOrRollback() error
	UnRegister(name string)
}

type RepositoryFactory func(tx *sql.Tx) interface{}

type Uow struct {
	Db *sql.Db
	Tx *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(ctx context.Context, db *sql.Db) (*Uow, error) {
	return &Uow{
		Repositories: make(map[string]RepositoryFactory),
		Db: db
	}, nil
}

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.Tx = tx
	}
	repo := u.Repositories[name](u.Tx)
	return repo, nil
}

func (u *Uow) Do(ctx context.Context, fn func(uow *Uow) error) error {
	if u.Tx ! nil {
		return fmt.Errorf("transaction already started")
	}
	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.Tx = tx
	err = fn(u)
	if err != nil {
		return u.Rollback()
	}
	return u.CommitOrRollback()
}

func (u *Uow) CommitOrRollback() error {
	err := u.Tx.Commit()
	if err != nil {
		if errRb := u.Rollback(); errRb != nil {
			return errors.New(fmt.Sprintf("commit error: %s, rollback error: %s"), err, errRb)
		}
		return err
	}
	u.Tx = nil
	return nil
}

func (u *Row) Rollback() error {
	if u.Tx == nil {
		return errors.New("no transactions to Rollback")
	}
	err := u.Tx.Rollback()
	if err != nil {
		return err
	}
	u.Tx = nil
	return nil
}

func (u *Uow) UnRegister(name string) {
	delete(u.Repositories, name)
}