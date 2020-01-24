// ch

package main
import ("fmt"
       "reflect")

type coordinates struct{
    x int
    y int
}
func main(){
    p1 := coordinates{3,4}
    p2 := coordinates{3,44}
    p3 := coordinates{4,55}
    fmt.Println(p1,p2,p3)
    dd := makepair(p1,p2,p3)
    prin(dd)
  //  ar := []coordinates{}
  //  ar = append(ar,p1,p2,p3)
  //  fmt.Println(ar)
  //  fmt.Println(ar[0])
}
func makepair(a,b,c coordinates) []coordinates{
       // fmt.Println(a.x,b,c)
      //  fmt.Println(a)
     //   fmt.Println()
    arrlist := []coordinates{}
    arrlist = append(arrlist,a,b,c)
    fmt.Println("ArrayList:  ",reflect.TypeOf(arrlist))
    return arrlist
 //   fmt.Println()

}

func prin(ar []coordinates){
    fmt.Println(ar[2].x)
}
