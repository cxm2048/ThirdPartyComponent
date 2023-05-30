package main

import (
    "context"
    "flag"
    "fmt"
    "github.com/go-redis/redis/v8"
    "log"
    "strconv"
)

type RedisConfig struct {
    Host     string
    Port     int
    Password string
    DB       int
    key      string
}

var (
    RedisCnf            RedisConfig
    gRedisClient        *redis.Client
    gRedisClusterClient *redis.ClusterClient
)

func init() {
    //flag.StringVar(&RedisCnf.Host, "redis_host", "172.22.223.47", "redis host")
    //flag.IntVar(&RedisCnf.Port, "redis_port", 6479, "redis port")
    //flag.StringVar(&RedisCnf.Password, "redis_password", "passwd", "redis password")
    flag.StringVar(&RedisCnf.Host, "redis_host", "192.168.11.135", "redis host")
    flag.IntVar(&RedisCnf.Port, "redis_port", 6379, "redis port")
    flag.StringVar(&RedisCnf.Password, "redis_password", "passwd", "redis password")
    flag.StringVar(&RedisCnf.key, "key", "uuid10", "redis test key")
}

func NewRedisClient() error {
    gRedisClient = redis.NewClient(&redis.Options{
        Addr:     RedisCnf.Host + ":" + strconv.Itoa(RedisCnf.Port),
        Password: RedisCnf.Password,
    })

    if gRedisClient == nil {
        log.Printf("redis connect failed\n")
        return fmt.Errorf("redis connect failed")
    }

    ret, err := gRedisClient.Ping(context.Background()).Result()
    if err != nil {
        log.Printf("ping ret:%v, err:%v\n", ret, err)
        return fmt.Errorf("ping failed")
    }

    return nil
}

func RedisSetNX() {

    ok1, err1 := gRedisClient.SetNX(context.Background(), "uuid3", "test", 0).Result()
    if err1 != nil {
        log.Printf("setnx err:%v\n", err1)
    }
    log.Printf("setnx ok:%v\n", ok1)

    ok2, err2 := gRedisClient.SetNX(context.Background(), "uuid3", "test", 0).Result()
    if err2 != nil {
        log.Printf("setnx err:%v\n", err2)
    }
    log.Printf("setnx ok:%v\n", ok2)
}

func NewRedisClusterClient() error {
    gRedisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
        Addrs: []string{
            RedisCnf.Host + ":" + strconv.Itoa(RedisCnf.Port),
        },
        Password: RedisCnf.Password,
    })
    log.Printf("gRedisClusterClient:%s\n", RedisCnf.Host+":"+strconv.Itoa(RedisCnf.Port))
    if gRedisClusterClient == nil {
        log.Printf("redis connect failed\n")
        return fmt.Errorf("redis connect failed")
    }

    ret, err := gRedisClusterClient.Ping(context.Background()).Result()
    if err != nil {
        log.Printf("ping ret:%v, err:%v\n", ret, err)
        return fmt.Errorf("ping failed")
    }

    return nil
}

func RedisClusterSetNX() {
    ok1, err1 := gRedisClusterClient.SetNX(context.Background(), RedisCnf.key, "test", 0).Result()
    if err1 != nil {
        log.Printf("setnx err:%v\n", err1)
    }
    log.Printf("setnx ok:%v\n", ok1)

    ok2, err2 := gRedisClusterClient.SetNX(context.Background(), RedisCnf.key, "test", 0).Result()
    if err2 != nil {
        log.Printf("setnx err:%v\n", err2)
    }
    log.Printf("setnx ok:%v\n", ok2)
}

func main() {
    //err := NewRedisClient()
    //if err != nil {
    //    log.Printf("RedisConnect failed, err:%v\n", err)
    //    return
    //}
    //
    //defer func() {
    //    err = gRedisClient.Close()
    //    if err != nil {
    //        log.Printf("redis close failed\n")
    //    }
    //}()
    //RedisSetNX()

    //ok, err := gRedisClient.HMSet(context.Background(), "test_key1", "key1", "value1", "key2", "value2", "key3", "value3").Result()
    //if err != nil || !ok {
    //    log.Printf("hmset err:%v, ok:%d\n", err, ok)
    //    return
    //}
    //
    //log.Printf("hmset ok:%v\n", ok)
    //
    //ret, err1 := gRedisClient.HMGet(context.Background(), "test_key3", "key1", "key2", "key3").Result()
    //if err1 != nil {
    //    log.Printf("hmget err:%v\n", err1)
    //    return
    //}
    //
    //log.Printf("hmget ret:%v\n", ret)

    err := NewRedisClusterClient()
    if err != nil {
        log.Printf("RedisConnect failed, err:%v\n", err)
        return
    }

    RedisClusterSetNX()
}
