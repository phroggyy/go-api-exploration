package main

import (
    "fmt"
    "os"
    "github.com/gin-gonic/gin"
    "net/http"
    "log"
    "github.com/phroggyy/go-api-exploration/models"
    "github.com/phroggyy/go-api-exploration/persistence"
)

func main() {
    port := "80"
    if len(os.Args) == 2 {
        port = os.Args[1]
    }
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
    
    api := router.Group("api");
    {
        api.GET("/", ctl.ApiIndex)
    }
    router.GET("/", ctl.Index)
    router.GET("/stream", ctl.Stream)
    fmt.Print("Routes initialised...\n")

    router.Static("/js", "./frontend/public/js")
    router.Static("/css", "./frontend/public/css")
    router.Static("/img", "./frontend/public/img")
    router.LoadHTMLFiles("frontend/public/index.html")

    // Seed our database with one user.
    db.DB.DropTableIfExists(&models.User{}).CreateTable(&models.User{})
    db.DB.Create(&models.User{Name:"John Doe",Email:"john.doe@example.com"})
    db.DB.Create(&models.User{Name:"Jane Doe",Email:"jane.doe@example.com"})

    // Start the application
    router.Run(":"+port)
}

func (ctl *Controller) ApiIndex(response *gin.Context) {
    users := models.Users{}
    ctl.session.DB.Find(&users)

    response.JSON(http.StatusOK, users)
}

func (ctl *Controller) Index(response *gin.Context) {
    response.HTML(http.StatusOK, "index.html", gin.H{})
}