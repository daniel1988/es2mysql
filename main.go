package main

import (
	"Common"
	"Store"
	"flag"
	"fmt"
)

var (
	flgEs        = flag.String("eshost", "127.0.0.1", "elastic search ip")
	flgEsPort    = flag.String("esport", "9200", "elastic search port")
	flgMysqlHost = flag.String("dbhost", "127.0.0.1", "mysql host")
	flgMysqlPort = flag.String("dbport", "3306", "mysql port")
	flgMysqlUser = flag.String("dbuser", "root", "mysql user")
	flgMysqlPwd  = flag.String("dbpwd", "root", "mysql password")
)

type Es2Mysql struct {
}

func NewEs2Mysql() *Es2Mysql {
	return &Es2Mysql{}
}

func (this *Es2Mysql) Mysql() {
	db, _ := Store.NewMysql("127.0.0.1", "3306", "root", "root", "dc_es_log", "utf8")
	fmt.Println(db)
	results := db.FetchAll("SELECT id,docid FROM t_es_202014 limit 10 ")
	for _, val := range results {
		fmt.Println(val)
	}
}

func (this *Es2Mysql) QueryEs() {

}

func main() {

	fmt.Println(*flgEs)
	fmt.Println("WeekTable: ", Common.WeekTable(1585733525, "t_es_base"))

	NewEs2Mysql().Mysql()
}
