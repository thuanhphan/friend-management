# friend-management-go
This is my first Go application

1. Install Go (the lastest stable version if you haven't already)
Then create a new Go project and initialize it with Go modules:
go mod init friend-management-go

2. Install chi for httprouter:


3. Install the migrate tool:
go get -u -d github.com/golang-migrate/migrate/v4/cmd/migrate
Then 
migrate create -ext sql -dir internal/db/migrations -seq create_tables

4. Install sqlboiler for generate model:
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
5. Create a sqlboiler.toml configuration file in your project root.
Then, generate the ORM code:
sqlboiler psql

6. Install config
go get github.com/kelseyhightower/envconfig