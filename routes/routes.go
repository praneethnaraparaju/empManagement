package routes

import (
    "empManagement/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.POST("/employees", controllers.CreateEmployee)
        api.GET("/employees", controllers.GetEmployees)
        api.GET("/employees/:user_id", controllers.GetEmployeeDetails)
        api.DELETE("/employees/:user_id", controllers.DeleteEmployee)
    }
}
