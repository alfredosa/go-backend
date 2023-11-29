# Golang backend template

to create migrations:

Next, create migration files using the following command:

```bash 

$ migrate create -ext sql -dir database/migrations/ -seq init_mg
```

You use -seq to generate a sequential version and init_mg is the name of the migration.

To migrate: 

```bash
$ migrate -path "./migrations/" -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

```
