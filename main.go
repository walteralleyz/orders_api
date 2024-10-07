package main

import (
	"net/http"
	"time"
	
	"orders-api/controller"
	
	"github.com/gin-gonic/gin"
	_ "github.com/dimiro1/banner/autoload"
)

func main() {
    var router = gin.Default()
    var controller = controller.New()
    
    var server = &http.Server{
        Addr: ":3000",
        Handler: router,
        ReadTimeout: 10* time.Second,
        WriteTimeout: 10* time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    
    router.GET("/:id", controller.GetCustomer)
    router.POST("/:name/:id", controller.PostCustomer)
    router.DELETE("/:id", controller.DeleteCustomer)
    router.PUT("/:id/:name", controller.UpdateCustomer)
    
    server.ListenAndServe()
}
