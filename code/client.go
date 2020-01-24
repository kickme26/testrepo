
// client request

package main
import (
    "fmt"
    "net/http"
    "io/ioutil"
    "net/url"
)

func main() {

  resp,err := http.PostForm("http://0.0.0.0:8080/a",url.Values{"Authkey":{"youareallowed12"},"username":{"guna"},"password":{"1234a"}})
 // resp,err := http.Get("http://0.0.0.0:8080/a" )
 // resp.Header.Set("Authkey","yourareallowed12")
  if err != nil{
    fmt.Println(err)
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil{
    fmt.Println(err)
  }

  fmt.Println(string(data))
}

