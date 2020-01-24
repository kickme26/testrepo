package main

import (
  "database/sql"
  "fmt"
  "os"
//  "log"
  "io"
  "encoding/csv"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "Hbpl@123"
  dbname   = "instagram"
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  //fmt.Println("Db:	",db)
  defer db.Close()

  err = db.Ping()
  fmt.Println("err:  ",err)
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
  csv_file,temper := os.Open("/home/guna/Downloads/data.csv")
  if temper != nil{
    panic(temper)
  }
  reading := csv.NewReader(csv_file)
  fmt.Println("reading done",reading)
  for {
    reading, err := db.Read()
    if err == io.EOF {
      break
    }
  fmt.Println(reading)
 // csv_file, ferr := os.Open("/home/guna/Downloads/md.csv")
  //  if ferr != nil {
    //  log.Fatal(ferr)
   // }
   // r := csv.NewReader(csv_file)
     // for {
//	record, err := r.Read()
      //  query := "insert into users(username,email,gender,country,ipaddress) values(?,?,?,?);"
//	if err == io.EOF {
//	  break
//	}
//	if err != nil {
//	  log.Fatal(err)
//	}
      //'  var username string
    //    err = db.QueryRow("select username from users where id = $1", 1).Scan(&username)
  //      if err != nil {
//	  log.Fatal(err)
      // }
    //    fmt.Println("name: ",username)'
   //     fmt.Println("___",record)
        }
}



