package utils

import (
    "database/sql"
    
    "orders-api/customer"
)

func FillCustomer(cs *customer.Customer, rows *sql.Rows) {
    for rows.Next() {
        rows.Scan(&cs.ID, &cs.Name)
    }
}
