package main

import (
    "Store"
    "Common"
    "fmt"
)

func main() {

    Common.Say()

    db, _ := Store.NewMysql("127.0.0.1", "3306", "root", "root", "dc_es_log", "utf8")
    fmt.Println(db)
    results := db.FetchAll("SELECT id,docid FROM t_es_202013 limit 10 " )
    for _,val := range results {
        fmt.Println(val)
    }

}