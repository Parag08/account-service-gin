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

- Clone the repository, install dependecies and build the binary

```bash
git clone https://github.com/Parag08/account-service-gin
cd account-service-gin
go get
go build -o bin/account-service-gin -v .
```

- Test the server locally

```bash
heroku local
```

- Create a heroku app

```bash
heroku create
```

- Add database plugin to heroku app

```bash
heroku addons:create heroku-postgresql:hobby-dev
```

- Use heroku config to get database url

```bash
heroku config
```

- Run migrations on the database url

```bash
migrate -source file://migrations -database <DATABASE_URL> up
```

- Open the app in browser

```bash
heroku open
```

## Test

## API Docs

## Reference

- https://github.com/gin-gonic/gin

