package main
import (
        "database/sql"
	"encoding/json"
        "fmt"
//        "github.com/jinzhu/gorm"
//	"net/http"
        "github.com/gin-gonic/gin"
        _ "github.com/lib/pq"
      )

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "Hbpl@123"
  dbname   = "instagram"
  )

func initdb () (*sql.DB){
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    fmt.Print(err)
  }
  err = db.Ping()
  if err != nil {
    fmt.Print(err)
  }
  return db
}
func getdetails (c *gin.Context){
  db := initdb()
  defer db.Close()
  var (
       name string
       mail string
      )
  id := c.Param("id")
  statement := fmt.Sprintf("select username, emailid from users where id = %s;", id)
  row := db.QueryRow(statement)
  err := row.Scan(&name, &mail)
 //       id := c.Param("id")
   //     row := db.QueryRow("select username,emailid  from users where id = ?;",id)
//      err = row.Scan(&name,&mail)
  if err  != nil {
     fmt.Println(err)
    }
  c.JSON(200, gin.H{
        "name":name,
        "mail":mail,
     })
}

type followersdetails struct{
	Id int       `json:"id"`
	Name string  `json:"name"`
}
type  Postfriends struct{
	Following int  `json: "following"`
	Followers int  `json: "followers"`
}


func friendrequest (c *gin.Context){
  db:= initdb()
  defer db.Close()
  var friends Postfriends
  c.Bind(&friends)
  query,err := db.Prepare("insert into friend(followers,following)values(?,?);")
  if err != nil{
    fmt.Println(err)
  }
  _,err = query.Exec(c.Following,c.Followers)
  if err != nil{
    fmt.Println(err)
  }
  c.JSON(201,gin.H{"followers":c.Followers,"following:"c.Following})
}

func getfollowers(c *gin.Context){
  db := initdb()
  defer db.Close()
  var (
//	F  followersdetails
	FL []followersdetails
  )

  id := c.Param("followers")
  statement := fmt.Sprintf("select f.followers,u.username from friends as f join users as u on f.followers = u.id where following = %s;",id)
  rows,err := db.Query(statement)
  if err != nil {
      fmt.Print(err)
  }
//  var F followersdetails
  for rows.Next() {
      var F followersdetails
      err = rows.Scan(&F.Id, &F.Name)
      fmt.Println("F-->",F.Id,F.Name)
      FL = append(FL, F)
      fmt.Println("FL:-->",FL)
      ba,err := json.Marshal(FL)
      if err != nil {
          fmt.Print(err)
       }
         }

   fmt.Println(">>",FL)
   c.JSON(200, gin.H{
                "count":len(FL),
		"result": FL})
//    FL = append(FL,F)
	//	fmt.Println(F.id,F.name)
//  answer := gin.H{
//		"result":FL,
//		"count":
}

func main() {
    ro := gin.Default()
    group_details := ro.Group("/persons")
      group_details.GET("user/:id", getdetails)
      group_details.POST("request/",friendrequest)
      group_details.GET("friendsof/:followers",getfollowers)
    ro.Run(":9000")

}

