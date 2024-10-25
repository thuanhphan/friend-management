# friend-management-go
This is my first Go application

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

-- Insert data into friendships table
INSERT INTO friendships (id, user_email, friend_email, status) VALUES <br />
(1, 'john.doe@example.com', 'jane.smith@example.com', 'friends'), <br />
(2, 'alice.johnson@example.com', 'bob.brown@example.com', 'subscribed'), <br />
(3, 'carol.white@example.com', 'john.doe@example.com', 'blocked'), <br />
(4, 'jane.smith@example.com', 'alice.johnson@example.com', 'friends'), <br />
(5, 'bob.brown@example.com', 'carol.white@example.com', 'subscribed'); <br />
