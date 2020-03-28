package Store

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
    Db *sql.DB
}

func NewMysql(host string, port string, user string, pwd string, dbname string, charset string) (*Mysql, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, pwd, host, port, dbname, charset)

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err.Error())
    }

    return &Mysql{db}, err
}

/**
 * 获取一行
 */
func (this *Mysql) FetchRow(sql string) map[string]string {

    rows, _ := this.Db.Query(sql)
    columns, _ := rows.Columns()

    scanArgs := make([]interface{}, len(columns))
    values := make([]interface{}, len(columns))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    rows.Next()
    //将行数据保存到record字典
    rows.Scan(scanArgs...)
    record := make(map[string]string)
    for i, col := range values {
        if col != nil {
            record[columns[i]] = string(col.([]byte))
        }
    }
    return record
}

/**
 * 获取所有数据
 */
func (this *Mysql) FetchAll(sql string) map[int]map[string]string {
    query, _ := this.Db.Query(sql)
    columns, _ := query.Columns()

    scanArgs := make([]interface{}, len(columns))
    values := make([][]byte, len(columns))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    results := make(map[int]map[string]string)
    i := 0

    for query.Next() {
        //将行数据保存到record字典
        if err := query.Scan(scanArgs...); err != nil {
            fmt.Println(err)
        }
        row := make(map[string]string)
        for k, v := range values {
            key := columns[k]
            row[key] = string(v)
        }
        results[i] = row
        i++
    }
    return results
}

func (this *Mysql) Execute(sql string, args ...interface{}) error {
    if len(args) == 1 {
        _, err := this.Db.Exec(sql)
        return err
    }

    stmt, err := this.Db.Prepare(sql)
    if err != nil {
        panic(err.Error())
    }
    _, err = stmt.Exec(args)
    if err != nil {
        panic(err.Error())
    }
    return nil
}

