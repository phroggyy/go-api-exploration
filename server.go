package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "log"
    "github.com/phroggyy/apitest/models"
    "github.com/phroggyy/apitest/persistence"
)

func main() {
    fmt.Print("Starting WebSockets...\n")
    go h.run()
    fmt.Print("Starting Gin...\n")
    router := gin.Default()

    db, err := persistence.Start()
    if err != nil {
        fmt.Print("Database connection error:", err)
        log.Fatal(err)
    }
    ctl := NewController(db)
    defer db.Close()
    fmt.Print("Connected to database...\n")

    // Set up our routes.
    // TODO: move this to a separate method
    router.GET("/", ctl.Index)
    router.GET("/stream", ctl.Stream)
    fmt.Print("Routes initialised...\n")

    // Seed our database with one user.
    db.DB.DropTableIfExists(&models.User{}).CreateTable(&models.User{})
    db.DB.Create(&models.User{Name:"John Doe",Email:"john.doe@example.com"})

    // Start the application
    router.Run(":8080")
}

func (ctl *Controller) Index(response *gin.Context) {
    users := models.Users{}
    ctl.session.DB.Find(&users)

    response.JSON(http.StatusOK, users)
}