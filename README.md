# friend-management-go
This is my first Go application - FM

### Run app
1. Clone src code <br />
2. Run docker-compose to migrate data: docker-compose up -d <br />
3. go run main/main.go

### Tech stack
1. Go & Chi
2. Postgresql
3. ORM: sqlboiler
4. Docker

### Setup
1. Install Go (the lastest stable version if you haven't already). Then create a new Go project and initialize it with Go modules: <br />
go mod init friend-management-go

2. Install chi for httprouter: go get -u github.com/go-chi/chi/v5

3. Install the migrate tool: <br />
go get -u -d github.com/golang-migrate/migrate/v4/cmd/migrate <br />
Then run the command below to generate 2 files sql to create table and drop table <br />
migrate create -ext sql -dir internal/db/migrations -seq create_tables
Notes: put the script into these files

4. Migrate data to docker container: <br />
docker-compose up -d
Notes: sometimes it requires run 2 times to make sure the container db.migrations-1 create tables

5. Install sqlboiler for generate model: <br />
go install github.com/volatiletech/sqlboiler/v4@latest <br />
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest <br />

6. Create a sqlboiler.toml configuration file in your project root. Then, generate the ORM code: <br />
sqlboiler psql

### Dummy data
-- Insert data into users table
INSERT INTO users (email, name) VALUES <br />
('john.doe@example.com', 'John Doe'), <br />
('jane.smith@example.com', 'Jane Smith'), <br />
('alice.johnson@example.com', 'Alice Johnson'), <br />
('bob.brown@example.com', 'Bob Brown'), <br />
('carol.white@example.com', 'Carol White'); <br />


