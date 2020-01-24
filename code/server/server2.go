//server 2 8081

package main
import (
  "github.com/gin-gonic/gin"
  "fmt"
)

func response(c *gin.Context) {
  c.String(200,"response from server2")
  fmt.Println("c.JSON(200,'response from server2')")

}

func server2() {
  host := "localhost"
  port := "8081"
  router := gin.New()
  router.Handle("GET", "/b", response)
  router.Run(host+ ":" +port)

}
func main() {
  server2()
}


