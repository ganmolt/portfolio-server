package main 

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

		router.LoadHTMLGlob("templates/*.html")

		data := "Hello Go/Gin!!"
    router.GET("/", func(c *gin.Context) {
				c.HTML(200, "index.html", gin.H{"data": data})
    })

    router.Run(":3001")
}
