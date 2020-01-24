package main
import (
	"github.com/gin-gonic/gin"
//	 "net/http"

	"fmt"
)
func resp(c *gin.Context){
	 fmt.Println("--->",c.FullPath(),c.Handler())
	 c.JSON(200,"helloooowwwwww")
}
func hello(c *gin.Context) {
	//fmt.Println(">>>>>",c.GetHeader)
	r := gin.New()
	d:=r.GET("/a",resp)
	fmt.Println(d)
//	r.ServeHTTP(c.Writer,c.Request)
	r.Run(":8001")
}
func presp(c* gin.Context){
        c.JSON(200,"Post")
    }


func main() {
	host := "localhost"
	port := "8000"
	ro := gin.New()
//	ro.Handle("GET","/a", hello)
//	req1:=ro.Handle("GET","http://0.0.0.0:8001/b",hello)
	//fmt.Println(req)
	//req1 := ro.Handle("POST","/b",hello)
	req2 := ro.Handle("POST","/a",presp)
	req1 := ro.Handle("GET","/b",resp)
	fmt.Println(req1,req2)
	ro.Run(host+":"+port)
}
