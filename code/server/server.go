package main
import (
	"github.com/gin-gonic/gin"
	//"fmt"
//	"net/http"
//	"io/ioutil"
)

func response(c *gin.Context){
  c.JSON(200,"you got a message!")
}
func server(){
  host:="localhost"
  port := "8080"
  router := gin.New()
  router.Handle("GET", "/a",response)
  router.Run(host+":"+port)

}
func main(){
  server()
}
