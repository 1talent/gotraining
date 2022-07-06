# Get started with golang

## cobra cli

## echo framework

## Run it

```bash
go build -o app
./app server
open http://localhost:1323
```

## Managing our golang models with postgres

## Part 1: Manage postgres schema with sql-migrate

```bash
go install -v github.com/rubenv/sql-migrate/...
```

## Part 2: SQL Boiler to generate golang models

```bash
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
```