# Account-service-gin

A account service written in gin to get to know gin better + try to create a robust way to store transactional data. Created using postgres databass.

## Installation

* Clone

```bash
git clone https://github.com/Parag08/account-service-gin
```

* Migrations

```bash
migrate -source file://migrations -database <DATABASE_URL> up
```

* Dependencies

```bash
go get
```

## Deployment

Huroku deployment steps

```bash
git clone https://github.com/Parag08/account-service-gin
cd account-service-gin
go get
go build -o bin/account-service-gin -v .
# this should start a local server
heroku local
heroku create

```

## Test

## API Docs

## Reference

- https://github.com/gin-gonic/gin
