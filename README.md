# friend-management-go
This is my first Go application

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 7343663 (Update docs, unit test)
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
<<<<<<< HEAD
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
<<<<<<< HEAD
INSERT INTO friendships (id, user_email, friend_email, status) VALUES
(1, 'john.doe@example.com', 'jane.smith@example.com', 'friends'),
(2, 'alice.johnson@example.com', 'bob.brown@example.com', 'subscribed'),
(3, 'carol.white@example.com', 'john.doe@example.com', 'blocked'),
(4, 'jane.smith@example.com', 'alice.johnson@example.com', 'friends'),
(5, 'bob.brown@example.com', 'carol.white@example.com', 'subscribed');
<<<<<<< HEAD

1. Install Go (the lastest stable version if you haven't already)
Then create a new Go project and initialize it with Go modules:
=======
1. Install Go (the lastest stable version if you haven't already) <br />
Then create a new Go project and initialize it with Go modules: <br />
>>>>>>> 9ff2063 (Edit README.md)
go mod init friend-management-go

2. Install chi for httprouter:

<<<<<<< HEAD
3. Install the migrate tool:
go get -u -d github.com/golang-migrate/migrate/v4/cmd/migrate
Then 
=======
=======
### 
1. Install Go (the lastest stable version if you haven't already). Then create a new Go project and initialize it with Go modules: <br />
go mod init friend-management-go

2. Install chi for httprouter: go get -u github.com/go-chi/chi/v5
>>>>>>> ba890b0 (Remove unused code)

3. Install the migrate tool: <br />
go get -u -d github.com/golang-migrate/migrate/v4/cmd/migrate <br />
Then <br />
>>>>>>> 9ff2063 (Edit README.md)
=======
>>>>>>> 7343663 (Update docs, unit test)
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

<<<<<<< HEAD
<<<<<<< HEAD
6. Install config
go get github.com/kelseyhightower/envconfig
=======
6. Install config <br />
go get github.com/kelseyhightower/envconfig
>>>>>>> 9ff2063 (Edit README.md)
=======
>>>>>>> ba890b0 (Remove unused code)
=======
>>>>>>> a78f4b5 (update script create table)
=======
INSERT INTO friendships (id, user_email, friend_email, status) VALUES <br />
(1, 'john.doe@example.com', 'jane.smith@example.com', 'friends'), <br />
(2, 'alice.johnson@example.com', 'bob.brown@example.com', 'subscribed'), <br />
(3, 'carol.white@example.com', 'john.doe@example.com', 'blocked'), <br />
(4, 'jane.smith@example.com', 'alice.johnson@example.com', 'friends'), <br />
(5, 'bob.brown@example.com', 'carol.white@example.com', 'subscribed'); <br />
>>>>>>> 07f2fdf (apply clean architecture)
