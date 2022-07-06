# Get started with golang

## cobra cli

## echo framework

## Run it

```bash
go build -o app
./app server
open http://localhost:1323
```

## Using golang cobra command line tool

```bash
go install github.com/spf13/cobra-cli@latest
```

This allows us to use the `cobra-cli` command in our command line.

## Managing our golang models with postgres

## Part 1: Manage postgres schema with sql-migrate

```bash
go install -v github.com/rubenv/sql-migrate/...
```

This allows us to use the `sql-migrate` command in our command line.

## Part 2: SQL Boiler to generate golang models

```bash
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
```