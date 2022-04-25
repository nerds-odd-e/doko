# TDD Service

## To Run the service:

### Using Makefile

#### Generate random key to decode signature


#### Create a bridge network locally, so app can talk to db:

```
$ make network
```

#### Create the databases `tdd` and `tdd-test`, run migrations against both databases, and seed main database.

```
$ make database
```

#### Start and run the application

```
$ make start
```

### Manual

#### Create a bridge network locally, so app can talk to db:

```
docker network create -d bridge tdd
```

#### For local database, can run:

The password below should be exported to env variable (or docker compose file) as DB_PASSWORD

```
docker run --rm --name test-postgres --net=cv -e POSTGRES_PASSWORD=<your password> -e POSTGRES_DB=access_control -p 5432:5432 -d postgres
```

Connect to the above DB and create a DB for test and [run](###To run migrations) migrations against both app DB and test DB:

```
docker run -e POSTGRES_PASSWORD=<your password> -p 5432:5432 postgres
create database access_control_test
```

#### To run migrations

We are using https://github.com/go-pg/pg for orm and https://github.com/go-pg/migrations for DB migrations.
Export DB_USER, DB_PASSWORD and DB_NAME run the following

```
cd migrations
go run *.go init
go run *.go up
go run *.go seed  #to get seed data
```


#### Running Test

### Memory Errors running test with race flag

Multiple memory related errors occur when running test with the race detection flag in OSX Monterey due to the design of automatic memory allocation in the OS.

To resolve, set environment vairable `MallocNanoZone=0` to disable OSX automatic memory allocation before running the test.

See [https://github.com/golang/go/issues/49138](https://github.com/golang/go/issues/49138)
