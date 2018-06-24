# Using database/sql

Here are some options one can choose to get started developing a Go DB application wtih TDD.

This project chose SQLite, with the go-sqlite3 driver, and the in-memory implementation for unit testing.

## In-memory DBs

Of the following, only RamSQL and SQLite supports the `database/sql` package.  RamSQL is designed for TDD.  SQLite can be used as both an in-memory and persistent store.

### [RamSQL](https://github.com/proullon/ramsql)

RamSQL has been written to be used in your project's test suite.

Unit testing in Go is simple, create a foo_test.go import testing and run go test ./.... But then there is SQL queries, constraints, CRUD...and suddenly you need a PostgresSQL, setup scripts and nothing is easy anymore.

The idea is to avoid setup, DBMS installation and credentials management as long as possible. A unique engine is tied to a single sql.DB with as much sql.Conn as needed providing a unique DataSourceName. Bottom line : One DataSourceName per test and you have full test isolation in no time.

### [SQLite](https://www.sqlite.org/inmemorydb.html)

This is broadly supportive of TDD via it's in-memory option, yet can also be used for "lite" durable storage requirements where a heavier DB isn't needed or desired.

There are quite a few `database/sql` drivers available for it.  The most popular is probably

[go-sqlite3](https://github.com/mattn/go-sqlite3)

### [go-memdb](https://github.com/hashicorp/go-memdb)

Provides the memdb package that implements a simple in-memory database built on immutable radix trees. The database provides Atomicity, Consistency and Isolation from ACID. Being that it is in-memory, it does not provide durability. The database is instantiated with a schema that specifies the tables and indices that exist and allows transactions to be executed.

## Mocking

### [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)

sqlmock is a mock library implementing sql/driver. Which has one and only purpose - to simulate any sql driver behavior in tests, without needing a real database connection. It helps to maintain correct TDD workflow.

## SQL Drivers

[SQL Drivers](https://github.com/golang/go/wiki/SQLDrivers)
