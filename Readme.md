# Pismo

## Start service

The services and their dependencies run within containers. To start the service, use the following command:

```console
$ docker-compose up
```

### Routes


#### Creating an account

```console
$ curl -XPOST -L "localhost:8080/accounts" -H "Content-Type: application/json" --data '{"document_number": "12345"}'
```

#### Getting an account by id

```console
$ curl -XGET -L "localhost:8080/accounts/{accounId}" -H "Content-Type: application/json"
```

#### Creating an account transaction

```console
$ curl -XPOST -L "localhost:8080/transactions" -H "Content-Type: application/json" --data '{"account_id": "{account_id}", "operation_type_id": 4, "amount": 100.30}'
```

## Run Tests

This project includes various types of tests, including acceptance tests that simulate the full cycle of a configured route. These tests will use an isolated database called test_pismo. To run the tests, use the following command:

```console
$ docker-compose --env-file ./env/test.env run tests
```

### Migrations

Migrations will automatically execute when running tests or spinning up the service. However, it's important to note that the current solution is not ideal for a production environment.

### Why GORM?

GORM is an Object-Relational Mapping (ORM) library that accelerates the development process. If, for some reason, GORM becomes the root cause of a performance issue, it can be replaced by modifying the domain package.

### Why httprouter?

`httprouter` is chosen because it is fast, similar to the standard library's router, and allows for segregating handlers by HTTP methods. Additionally, it implements OPTIONS handling for free.

### Why not Gin?

While Gin is a web framework that offers a wide range of features such as middleware, CORS support, JSON responses, and more, this project aims to balance service performance by using the standard library as much as possible while integrating with GORM.
