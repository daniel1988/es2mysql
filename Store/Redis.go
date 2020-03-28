package Store

import (
    // "fmt"
    Common "Common"
    "github.com/garyburd/redigo/redis"
    "time"
)

type Redis struct {
    pool *redis.Pool
}

// 使用redis连接池，用前Get，用完Close
func NewRedisPool(addr string, nrDb int) *Redis {
    defer Common.CheckPanic()

    pool := &redis.Pool{
        MaxIdle:     8,
        IdleTimeout: time.Minute,
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", addr)
            if err != nil {
                return nil, err
            }

            _, err = c.Do("SELECT", nrDb)
            return c, err
        },
    }

    return &Redis{pool}
}

// redis的incr操作是原子性递增的数字，可以用来生成msgid
func (this *Redis) Incr(key string) (int64, error) {
    defer Common.CheckPanic()
    rc := this.pool.Get()
    defer rc.Close()

    return redis.Int64(rc.Do("INCR", key))
}

// 标识请求的哈希值和消息ID的映射关系
func (this *Redis) SetEx(key, value string, ttl int64) {
    defer Common.CheckPanic()
    rc := this.pool.Get()
    defer rc.Close()

    rc.Do("SETEX", key, ttl, value)
}

func (this *Redis) Set(key, value string) {
    defer Common.CheckPanic()
    rc := this.pool.Get()
    defer rc.Close()

    rc.Do("SET", key, value)
}

// 根据哈希值查询消息ID
func (this *Redis) Get(key string) (string, error) {
    if len(key) == 0 {
        return "", nil
    }

    defer Common.CheckPanic()
    rc := this.pool.Get()
    defer rc.Close()

    res, err := rc.Do("GET", key)
    if err != nil {
        return "", err
    }
    if res == nil {
        return "", nil
    }

    return redis.String(res, nil)
}

// 广播队列，使用列表LIST
// 新消息广播队列, 返回队列长度
func (this *Redis) Rpush(queue string, msg []byte) (int64, error) {
    if len(queue) == 0 || len(msg) == 0 {
        return 0, nil
    }

    defer Common.CheckPanic()
    rc := this.pool.Get()
    defer rc.Close()

    return redis.Int64(rc.Do("RPUSH", queue, msg))
}

// 从消息广播队列取一个消息，非阻塞，需要循环调用
func (this *Redis) Lpop(queue string) ([]byte, error) {
    if len(queue) == 0 {
        return nil, nil
    }

    defer Common.CheckPanic()
    rc := this.pool.Get()
    defer rc.Close()

    // 返回ErrNil是因为队列空，算不上错误
    msg, err := redis.Bytes(rc.Do("LPOP", queue))
    if err == redis.ErrNil {
        return nil, nil
    }

    return msg, err
}

