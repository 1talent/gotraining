package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TxFn func(boil.ContextExecutor) error

func WithTransaction(ctx context.Context, db *sql.DB, fn TxFn) error {
	return WithConfiguredTransaction(ctx, db, nil, fn)
}

func WithConfiguredTransaction(ctx context.Context, db *sql.DB, options *sql.TxOptions, fn TxFn) error {
	tx, err := db.BeginTx(ctx, options)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if txErr := tx.Rollback(); txErr != nil {
				fmt.Println("Failed to roll back transaction after recovering from panic")
			}
			panic(p)
		} else if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				fmt.Println("Failed to roll back transaction after receiving error")
			}
		} else {
			err = tx.Commit()
			if err != nil {
				fmt.Println("Failed to commit transaction")
			}
		}
	}()

	err = fn(tx)
	return err
}
