/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/1talent/gotraining/internal/config"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Executes all migrations which are not yet applied.",
	Long:  `This is our custom database migration command that allows us to run a series of golang functions to load sql files and execute against our postgresql database.`,
	Run:   migrateCmdFunc,
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	migrate.SetTable(config.DatabaseMigrationTable)
}

func migrateCmdFunc(cmd *cobra.Command, args []string) {
	n, err := applyMigrations()
	if err != nil {
		fmt.Printf("Error while applying migrations: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Applied %d migrations.\n", n)
}

func applyMigrations() (int, error) {
	ctx := context.Background()
	serviceConfig := config.DefaultServiceConfigFromEnv()
	fmt.Println("DatabaseMigrationFolder", config.DatabaseMigrationFolder)
	db, err := sql.Open("postgres", serviceConfig.Database.ConnectionString())
	if err != nil {
		return 0, err
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		return 0, err
	}

	// In case an old default sql-migrate migration table (named "gorp_migrations") still exists we rename it to the new name equivalent
	// in sync with the settings in dbconfig.yml and config.DatabaseMigrationTable.
	if _, err := db.Exec(fmt.Sprintf("ALTER TABLE IF EXISTS gorp_migrations RENAME TO %s;", config.DatabaseMigrationTable)); err != nil {
		return 0, err
	}
	migrations := &migrate.FileMigrationSource{
		Dir: config.DatabaseMigrationFolder,
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return 0, err
	}

	return n, nil
}
