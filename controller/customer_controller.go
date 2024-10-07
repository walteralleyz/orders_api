package controller

import (
    "net/http"
	"log"
	
	"orders-api/customer"
	"orders-api/repository"
	"orders-api/utils"
	
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
    r *repository.CustomerRepository
}

func New() CustomerController {
    cc := CustomerController{
        r: &repository.CustomerRepository{},
    }
    
    cc.r.CreateTable()
    
    return cc
}

func (cc *CustomerController) GetCustomer(c *gin.Context) {
    var cs customer.Customer
    
    cc.r.LoadContext(c)
    
    c.ShouldBindUri(&cs)
    
    utils.FillCustomer(&cs, cc.r.Retrieve(cs.ID))
    
    if cs.Name != "" {
        c.JSON(http.StatusOK, gin.H{
            "msg": "hello",
            "name": cs.Name,
            "id": cs.ID,
        })
        
        return
    }
    
    utils.TreatErrorResponse(c, "Error querying for customer id")
}

func (cc *CustomerController) PostCustomer(c *gin.Context) {
    var cs customer.Customer
    
    cc.r.LoadContext(c)
    
    c.ShouldBindUri(&cs)
    
    cc.r.Create(cs.ID, cs.Name)
    
    log.Println("Customer registered into db!")
    
    c.JSON(http.StatusCreated, gin.H{"msg": "created!"})
}

func (cc *CustomerController) DeleteCustomer(c *gin.Context) {
    var cs customer.Customer
    
    cc.r.LoadContext(c)
    
    c.ShouldBindUri(&cs)
    
    utils.FillCustomer(&cs, cc.r.Retrieve(cs.ID))
    
    if cs.Name == "" {
        utils.TreatErrorResponse(c, "Customer id not found!")
        return
    }
    
    cc.r.Remove(cs.ID)
    
    log.Println("Customer deleted!")
    
    c.JSON(http.StatusAccepted, gin.H{"msg": "deleted!"})
}

func (cc *CustomerController) UpdateCustomer(c *gin.Context) {
    var cs customer.Customer
    var temp customer.Customer
    
    cc.r.LoadContext(c)
    
    c.ShouldBindUri(&cs)
    
    utils.FillCustomer(&temp, cc.r.Retrieve(cs.ID))
    
    if temp.Name == "" {
        utils.TreatErrorResponse(c, "Customer id not found!")
        return
    }
    
    cc.r.Update(cs.Name, cs.ID)
    
    log.Println("Customer udpated")
    
    c.JSON(http.StatusAccepted, gin.H{"msg": "Customer updated!"})
}

