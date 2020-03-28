package main

import (
    "Store"
    "fmt"
)

func main() {

    redis := Store.NewRedisPool("127.0.0.1:6379", 0)
    fmt.Println(redis.Incr("foo"))

}