package cmd

import (
    "github.com/gin-gonic/gin"
    "github.com/JuacCrack/GOapi/config"
    "github.com/JuacCrack/GOapi/routes"
    swaggerFiles "github.com/swaggo/files" 
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "github.com/JuacCrack/GOapi/docs" 
    "log"
)

func Execute() {
    config.LoadConfig()
    r := gin.Default()

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    routes.RegisterRoutes(r)
    if err := r.Run(); err != nil {
        log.Fatalf("Error starting server: %s", err)
    }
}
