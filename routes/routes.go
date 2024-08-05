package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/JuacCrack/GOapi/controllers"
)

func RegisterRoutes(r *gin.Engine) {
    r.GET("/contacts", controllers.GetContacts)
    r.GET("/contacts/:id", controllers.GetContact)
    r.POST("/contacts", controllers.CreateContact)
    r.PUT("/contacts/:id", controllers.UpdateContact)
    r.DELETE("/contacts/:id", controllers.DeleteContact)
}
