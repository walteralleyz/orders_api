## ORDERS API
#### Some simple CRUD API in Golang

This repository is part of my learning journey in golang. <br />
It's made with pure golang, Gin framework and SQLITE3. <br />
It does not follow the HTTP pattern, every command is made by url path. <br />

### Specs
Go Version  : 1.23 <br />
Default port: 3000 <br />

### Requisites
Golang  installed <br />
SQLITE3 installed <br />

### Routes
<mark>Retrieve</mark> GET    /:id       <br />
<mark>Create</mark>   POST   /:name/:id <br />
<mark>Remove</mark>   DELETE /:id       <br />
<mark>Update</mark>   PUT    /:id/:name <br />

### Hierarchie
<pre>
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
</pre>
