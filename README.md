# Simple-Bank
This is a bank service where we will provide API to fronted for doing things such as:

1. Create and manage bank accounts, which are composed of owner’s name, balance, and currency.
2. Record all balance changes to each of the account. So every time some money is added to or subtracted from the account, an account entry record will be created.
3. Perform a money transfer between 2 accounts. This should happen within a transaction, so that either both accounts’ balance are updated successfully or none of them are.

## Setup local development
### Install tools
- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
    ```bash
    brew install golang-migrate
    ```
- [Sqlc](https://github.com/kyleconroy/sqlc#installation)
    ```bash
    brew install sqlc
    ```

## Setup infrastructure
- Start postgres container:
    ```bash
    make postgres
    ```
- Create simple_bank database:
    ```bash
    make createdb
    ```
- Run db migration up all versions:
    ```bash
    make migrateup
    ```
- Run db migration down all versions:
    ```bash
    make migratedown
    ```

## How to generate code
- Generate SQL CRUD with sqlc:
    ```bash
    make sqlc
    ```
- Create a new db migration:
    ```bash
    migrate create -ext sql -dir db/migration -seq <migration_name>
    ```

## How to run
- Run test:
    ```bash
    make test
    ```