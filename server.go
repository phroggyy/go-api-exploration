package main

import (
    "fmt"
    "flag"
    "strings"
    "os"
    "github.com/gin-gonic/gin"
    "net/http"
    "log"
    "github.com/phroggyy/go-api-exploration/models"
    "github.com/phroggyy/go-api-exploration/persistence"
)

var siteName string

func main() {
    // Declare our flags
    var safeMode bool
    var certFile string
    var keyFile  string
    var port     string
    flag.BoolVar(&safeMode, "safe", false, "Sets whether the application should be served over https or not")
    flag.StringVar(&certFile, "cert", "./cert.pem", "Set the certificate file to be used")
    flag.StringVar(&keyFile, "key", "./privkey.pem", "Set the private key to be used")
    flag.StringVar(&port, "p", "80", "Set the port to run on")
    flag.Parse()
    fmt.Println("cert is " + certFile)

    // If the user left the cert and keyfiles default, ask if they're sure...
    if safeMode && certFile == "./cert.pem" && keyFile == "./privkey.pem" {
        fmt.Println("You didn't provide a location for your certificate and/or private key. We will try to use \"./cert/pem\" and \"./privkey.pem\" respectively.\nAre you sure you want to continue?\n")
        var answer string
        certain := false
        for ;!certain; {
            fmt.Scanln(&answer)
            answer = strings.ToLower(answer)
            if answer == "y" || answer == "yes" {
                certain = true
            } else if answer == "n" || answer == "no" {
                certain = false
                os.Exit(0)
            } else {
                fmt.Println("You did not provide a valid answer. Please try again.\n")
            }
        }
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

    // Seed our database with two users.
    db.DB.DropTableIfExists(&models.User{}).CreateTable(&models.User{})
    db.DB.Create(&models.User{Name:"John Doe",Email:"john.doe@example.com"})
    db.DB.Create(&models.User{Name:"Jane Doe",Email:"jane.doe@example.com"})

    if safeMode {
        // Start the application on HTTP in a goroutine
        go http.ListenAndServe(":80", http.HandlerFunc(redir))
        // Only respond to https
        router.Run(":443")
    } else {
        // Start the application
        router.Run(":"+port)
    }
}

func redir(w http.ResponseWriter, req *http.Request) {
    http.Redirect(w, req, "https://"+siteName+req.RequestURI, http.StatusMovedPermanently)
}

func (ctl *Controller) ApiIndex(response *gin.Context) {
    users := models.Users{}
    ctl.session.DB.Find(&users)

    response.JSON(http.StatusOK, users)
}

func (ctl *Controller) Index(response *gin.Context) {
    response.HTML(http.StatusOK, "index.html", gin.H{})
}