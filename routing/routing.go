package routing

func Setup(router *gin.Engine) {
    *router.GET("/", Index)
}