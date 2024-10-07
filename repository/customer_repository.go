package repository

import (
    "database/sql"
    "log"
    
    "orders-api/utils"
    
    _"github.com/mattn/go-sqlite3"
    "github.com/gin-gonic/gin"
)

type Repository interface {
    CreateTable()
    Retrieve() *sql.Rows
    Create()
    Remove()
    Update()
    LoadContext(c *gin.Context)
}

type CustomerRepository struct {
    db *sql.DB
    context *gin.Context
}

func (cr *CustomerRepository) CreateTable() {
    db, err := sql.Open("sqlite3", "customers.db")

    cr.db = db

    st, err := cr.db.Prepare(`
    CREATE TABLE IF NOT EXISTS
        customers (
            id INTEGER PRIMARY KEY,
            name VARCHAR(255)
        )
    `)
    
    if err != nil {
        log.Println("Error creating table")
        return
    }
    
    st.Exec()
    
    log.Println("table created")
}

func (cr *CustomerRepository) Retrieve(id string) *sql.Rows {
    rows, err := cr.db.Query(`
    SELECT id, name
    FROM customers
    WHERE id = 
    ` + id)
    
    if err != nil {
        utils.TreatErrorResponse(cr.context, "Error querying for customer id")
        return nil
    }
    
    return rows
}

func (cr *CustomerRepository) Remove(id string) {
    st, err := cr.db.Prepare(`
    DELETE FROM customers
    WHERE id = ?
    `)
    
    st.Exec(id)
    
    if err != nil {
        utils.TreatErrorResponse(cr.context, "error removing record")
        return
    }
}

func (cr *CustomerRepository) Update(id string, name string) {
    st, err := cr.db.Prepare(`
    UPDATE customers
    SET name = ?
    WHERE id = ?
    `)
    
    st.Exec(name, id)
    
    if err != nil {
        utils.TreatErrorResponse(cr.context, "error updating record")
        return
    }
}

func (cr *CustomerRepository) Create(id string, name string) {
    st, err := cr.db.Prepare(`
    INSERT INTO customers (id, name)
    VALUES (?, ?)
    `)
    
    st.Exec(id, name)
    
    if err != nil {
        utils.TreatErrorResponse(cr.context, "error inserting record")
        return
    }
}

func (cr *CustomerRepository) LoadContext(c *gin.Context) {
    cr.context = c
}
