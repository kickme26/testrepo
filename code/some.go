package main
import (
        "github.com/gin-gonic/gin"
        "fmt"
	//"net"
	"io/ioutil"
	"net/http"
//	"os/exec"

)

//func makerequest(){
//	cmd := exec.Command("GET http://localhost:8000/a")
//	out,err := cmd.Run().Output()
//	fmt.Println("output:", out)

//}

func dorequest(){
	resp,err := http.Get("http://localhost:8001/a")
	if err != nil {
	  fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
func response(c *gin.Context){
         fmt.Println("--->",c.FullPath(),c.Handler())
         c.JSON(200,"you got a message!")
	 //ip := c.ClientIP()
         //fmt.Println(">>>>>>>",string(g))
}
func getresponse(c *gin.Context) {
        router2 := gin.New()
        router2.GET("/a",response)
	fmt.Println("____2____")
        router2.Run(":8001")
	//dorequest()
	//makerequest()
}
func postresponse(c* gin.Context){
        c.JSON(200,"Post")
    }


func main() {
        host := "localhost"
        port := "8000"
        router1 := gin.New()
       // req1 := ro.Handle("POST","/a",presp)
        router1.Handle("GET","/b",getresponse)
       //ro.Run("localhost:8000")
        router1.Run(host+":"+port)
	fmt.Println("___1__")
	resp,err := http.Get("http://localhost:8000/b")
        if err != nil {
          fmt.Println(err)
        }
        body, err := ioutil.ReadAll(resp.Body)
        fmt.Println(string(body))


}
