package repository

import (
    "testing"

    "orders-api/customer"
)

func TestRetrieve(t *testing.T) {
    var temp customer.Customer
    var want = "test"

    cr := CustomerRepository{}
    cr.CreateTable()
    
    cr.Create("100", want)
    
    rows := cr.Retrieve("100")
    
    for rows.Next() {
        rows.Scan(&temp.ID, &temp.Name)
    }
    
    if temp.Name != want {
        t.Errorf("got %s, wanted %s", temp.Name, want)
    }
}
