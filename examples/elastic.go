package main

import (
    "Store"
    "fmt"
)

func main() {

    redis := Store.NewStaticStore("10.105.249.250", 9200)
    fmt.Println(redis.Incr("foo"))

}