## ORDERS API
#### Some simple CRUD API in Golang

This repository is part of my learning journey in golang.
It's made with pure golang, Gin framework and SQLITE3.
It does not follow the HTTP pattern, every command is made by url path.

### Specs
Go Version  : 1.23
Default port: 3000

### Requisites
Golang  installed
SQLITE3 installed

### Routes
<mark>Retrieve</mark> GET    /:id
<mark>Create</mark>   POST   /:name/:id
<mark>Remove</mark>   DELETE /:id
<mark>Update</mark>   PUT    /:id/:name

### Hierarchie
├── *banner.txt*
├── controller
│   └── customer_controller.go
├── customer
│   └── customer.go
├── customers.db
├── go.mod
├── go.sum
├── **main.go**
├── orders-api
├── README.md
├── repository
│   ├── customer_repository.go
│   └── customer_repository_test.go
└── utils
    ├── customer_utils.go
    └── error_utils.go

