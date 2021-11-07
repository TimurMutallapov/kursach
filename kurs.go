package main
import (
    "time"
    "database/sql"
    "fmt"
    "net/http"
    "net"
    "github.com/likexian/whois-go"

      _ "github.com/mattn/go-sqlite3"
)

func main() {

    db, err := sql.Open("sqlite3", "far.db")
        if err != nil {
            panic(err)
          }


    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        http.ServeFile(w, r, "ggwp.html")
          })



    http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request){

        name := r.FormValue("username")
        result, err := whois.Whois(name)

        if err != nil {
                fmt.Println(err)
        }



        fmt.Println(result)
        fmt.Fprintf(w, "Whois : %s", result)

        addr,err := net.LookupIP(name)
        fmt.Println(addr)

        if err != nil{
       fmt.Println(err)
      }

      str:= fmt.Sprintf("%s", addr)
      fmt.Println(str)



        createdAt := time.Now()


        resultdb,err := db.Exec("insert into users (TimeRequest,Domain, IP) values ($1, $2,$3 )",
        createdAt, name, str)
        if err != nil{
            panic(err)
        }
        fmt.Println(resultdb.LastInsertId())  // id последнего добавленного объекта
        //fmt.Println(result.RowsAffected())  // количество добавленных строк
    })
    //--------------------------------------------------------------------------
    //defer db.Close()


    fmt.Println("Server is listening...")
    http.ListenAndServe(":8181", nil)
}
