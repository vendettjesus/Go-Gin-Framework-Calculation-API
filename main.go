package main

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.GET("/calcularCalificacion", func(c *gin.Context) {
    name := c.DefaultQuery("name", "")
    cal1, _ := strconv.ParseFloat(c.DefaultQuery("cal1", "0"), 64)
    cal2, _ := strconv.ParseFloat(c.DefaultQuery("cal2", "0"), 64)
    cal3, _ := strconv.ParseFloat(c.DefaultQuery("cal3", "0"), 64)
   
    promedio := (cal1 + cal2 + cal3) / 3.0

    c.JSON(http.StatusOK, gin.H{
        "name":    name,
        "promedio": promedio,
    })
})

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}