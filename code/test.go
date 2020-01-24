// last updated 23/12/2019   1pm

package main
import ("fmt"
        "strings"
        "math"
        "sort"
       )

func strConversion(pt []int)string{
    newarr := []string{}
    for i:=range pt{
        newarr = append(newarr,fmt.Sprint(pt[i]))
    }
    conv_ar := strings.Join(newarr,",")
    return conv_ar
}

func generateKey(pts []int, ma map[string]int)(bool, string,[]int){
    copy_pts := make([]int,len(pts))
    copy(copy_pts,pts)
    sort.Ints(pts)
    converted_arr := strConversion(pts)
    if _,yes :=ma[converted_arr];yes{
         return true,converted_arr,copy_pts
    }
    return false,converted_arr,copy_pts
}
func  calculateDistance(first,second Coordinates)float64{
    dx := second.X - first.X  // x2-x1
    dy := second.Y - first.Y  // y2-y1
    distance := math.Sqrt(float64(dx*dx + dy*dy))
    return distance
}
func calculateDiagonal(l,h float64)float64{
    diagonal := math.Sqrt(float64(h*h + l*l))
   // fmt.Println("--------:",chk)
    return diagonal
}



func getcount(arr []Coordinates)int{
  //  fmt.Println("the points:",arr)
  rectpoints := make(map[string]int)
  for apts:=0 ; apts<len(arr);apts++{
    for bpts:=0;bpts<len(arr);bpts++{
      for cpts:=0;cpts<len(arr);cpts++{
        for dpts:=0;dpts<len(arr);dpts++{
          if dpts != apts && cpts != dpts && bpts != dpts && apts != bpts && bpts != cpts && apts != cpts{
            allpts := []int{arr[apts].X,arr[apts].Y,arr[bpts].X,arr[bpts].Y,
            arr[cpts].X,arr[cpts].Y,arr[dpts].X,arr[dpts].Y}
            checkkey,string_pairs,vertices := generateKey(allpts[:],rectpoints)
            if checkkey{
              dpts++
              continue
            }
            listalldistance := getDistance(arr[apts],arr[bpts],arr[cpts],arr[dpts])
            validateDistance(listalldistance [:], rectpoints, string_pairs,vertices)
          }// fr dp
        }//if cp
      }//for cp 
    }//fr bp
  }//fr ap
  return len(rectpoints)
}
func getDistance(a1,b1,c1,d1 Coordinates)[]float64{
    alldist := []float64{}
    ab := calculateDistance(a1,b1)
    bc := calculateDistance(b1,c1)
    cd := calculateDistance(c1,d1)
    da := calculateDistance(d1,a1)
    ac := calculateDistance(a1,c1)
 //   fmt.Println(">ab>bc>cd",ab,bc,cd,da,ac)
 // append in a slice and sorting it
    alldist = append(alldist,ab,bc,cd,da,ac)
    return alldist
}
func validateDistance(dlist []float64, rct_points map[string]int,j_s string,rectangle_vertices[]int) bool{
    sort.Float64s(dlist)
    if dlist[0] == dlist[1] && dlist[2]==dlist[3] && dlist[0]<dlist[2] && dlist[2]<dlist[4]{
        length := dlist [2]
        height := dlist[0]
        diag := dlist[4]
        lh2 := calculateDiagonal(length,height)
     // fmt.Println("length",lh2, "ac^2 = ",diag)
        if lh2==diag{
            rct_points[j_s] = 1
            fmt.Println("vertices :",rectangle_vertices)
            fmt.Println("map: ",rct_points)
            }
       }
    return false
}
type Coordinates struct {
    X int
    Y int
}

//func makepair(v1,v2,v3,v4 Coordinates)[]Coordinates{
//    arlist := []Coordinates{}
//    arlist = append(arlist,v1,v2,v3,v4)
//    return arlist
//}

func main(){
   // ar2 := arr[][2]int{{8,3},{0,-1},{-2,3},{6,7}}
   // pts := Matrix{{-3, 2}, {-3, -3}, {4, -3}, {4, 3}}
    c1 := Coordinates{-3,-3}
    c2 := Coordinates{4,2}
    c3 := Coordinates{6,7}
    c4 := Coordinates{-3,2}
    c5 := Coordinates{0,-1}
    c6 := Coordinates{-2,3}
    c7 := Coordinates{8,3}
    c8 := Coordinates{4,-3}
    arrayofpoints := []Coordinates{}
    arrayofpoints = append(arrayofpoints,c1,c2,c3,c4,c5,c6,c7,c8)
    // arrayofpoints:= makepair(c1,c2,c3,c4)

    // pts2 := Matrix{{-3,2} ,{4,2},{4,-3},{-3,-3},{8,3},{0,-1},{-2,3},{6,7}}
    // ar1 := Matrix{{-3,2} ,{4,2},{4,-3},{-3,-3},{8,3},{0,-1},{-2,3},{6,7}}
    // makematrix(ar1)
    //fmt.Println("\nThe total rectangle formed by 2 is:",ar1)
    // fmt.Println("type--",pts)
    check_pts  :=  getcount(arrayofpoints)
    // check_pts2 :=  getcount(pts2)
    fmt.Println("\n\nTotal Rectangles Formed:      ",check_pts)
    // fmt.Println(">Total rectangles formed:>",check_pts2)

}









