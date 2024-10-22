# friend-management-go
This is my first Go application

### 
1. Install Go (the lastest stable version if you haven't already). Then create a new Go project and initialize it with Go modules: <br />
go mod init friend-management-go

2. Install chi for httprouter: go get -u github.com/go-chi/chi/v5

3. Install the migrate tool: <br />
go get -u -d github.com/golang-migrate/migrate/v4/cmd/migrate <br />
Then <br />
migrate create -ext sql -dir internal/db/migrations -seq create_tables

4. Install sqlboiler for generate model: <br />
go install github.com/volatiletech/sqlboiler/v4@latest <br />
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest <br />

5. Create a sqlboiler.toml configuration file in your project root. Then, generate the ORM code: <br />
sqlboiler psql

