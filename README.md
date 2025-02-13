# Gatorblog
Gatorblog (or gator) is an RSS feed CLI tool for collecting and browsing posts, written in Go for the boot.dev curriculum.

## Requirements
### Running
* Golang, version 1.23+
  * You can use the instructions from [boot.dev](https://github.com/bootdotdev/bootdev) if needed.
  * Easy way is:
    * `curl -sS https://webi.sh/golang | sh`
  * Run `go version` to confirm after installation.
* Postgres
  * For Linux:
    * `sudo apt update`
    * `sudo apt install postgresql postgresql-contrib`
    * check version with `psql --version`
    * update password with `sudo passwd postgres`
      * I used 'postgres' for my local password.
      * Update the connection string in the Setup config file if you use something different.
    * start service with `sudo service postgresql start`
    * launch psql cli `sudo -u postgres psql`
    * `CREATE DATABASE gator;`
    * `\c gator`
    * `ALTER USER postgres PASSWORD '<password you set>';`
    * `SELECT version();` to confirm setup
    * Make sure you end any psql queries in the CLI with ;
    * Use `\q` to quit the psql cli.
  * For Mac:
    * `brew install postgresql@15`
    * password update not necessary.
    * `brew services start postgresql`
    * `psql postgres`
    * create database as above.

### Developing
* Goose, a db migration tool
  * `go install github.com/pressly/goose/v3/cmd/goose@latest`
  * run `goose -version` to confirm installation.
  * in `sql/schema`, use `goose postgres <connection_string> up/down` to migrate up or down.
    * connection string is the db_url.
* SQLC, a tool to generate Go code from SQL queries
  * `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
  * run `sqlc version` to make sure it's installed correctly.
  * in project root, run `sqlc generate` to generate Go code based on sql queries.
    * default location is `sql/queries`
    * can edit sqlc.yaml if desired.

## Setup
For Unix systems, create a file "~/.gatorconfig.json" with the following fields:
* db_url: postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
* current_user_name: <you>

If you named it or placed it anywhere differently, update the Read() function in the internal/config directory to read from the correct location.

In project root, run `go install` to install all required modules. From there, use `./gatorblog` to run commands, such as `./gatorblog register bryce`.