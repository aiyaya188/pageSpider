package cache

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"../common"
	"../database"
	"github.com/garyburd/redigo/redis"
)

var redisconn redis.Conn
var (
	Pool *redis.Pool
)

func init() {
	/*
		redisHost := common.GetConfig("redis", "address").String()
		Pool = NewRedisPool(redisHost)
		cleanupHook()
	*/
}
func cleanupHook() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

//var OffsetMap map[string]string
func NewRedisPool(server string) *redis.Pool {
	idelTimeout := 240 * time.Second
	// **重要** 设置读写超时

	readTimeout := redis.DialReadTimeout(time.Second * time.Duration(5))
	writeTimeout := redis.DialWriteTimeout(time.Second * time.Duration(5))
	conTimeout := redis.DialConnectTimeout(time.Second * time.Duration(5))
	redisPool := &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:   5,
		MaxActive: 0,
		// **重要** 如果空闲列表中没有可用的连接
		// 且当前Active连接数 < MaxActive
		// 则等待
		Wait:        true,
		IdleTimeout: idelTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server,
				readTimeout, writeTimeout, conTimeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	return redisPool
}

func initCache() {
	if redisconn == nil {
		conn, err := redis.Dial("tcp", common.GetConfig("redis", "address").String())
		if err == nil {
			redisconn = conn
			/*redisconn.Do("SELECT", 1)*/
		} else {
			fmt.Println("redis connect fail!")
		}
	}
}
func PutCache(key string, value string, timeout int) {
	/*
		initCache()
		_, err := redisconn.Do("SET", key, value, "EX", timeout)

		if err != nil {
			fmt.Println(err)
		}
	*/
	//database.OffsetMap[key]=value
	database.OffsetMap.Store(key, value)
}
func GetCache(key string) string {
	/*
		initCache()
		value, err := redis.String(redisconn.Do("GET", key))
		if err != nil {
			return ""
		}

		return value
	*/
	// return database.OffsetMap[key]
	res, ok := database.OffsetMap.Load(key)
	if ok == false {
		return ""
	}
	return res.(string)
}

func DelCache(key string) string {
	/*if GetConfig("system", "usecache").String() != "true" {*/
	//fmt.Println("disable use cache")
	//return ""
	/*}*/
	/*
		initCache()
		_, err := redis.String(redisconn.Do("DEL", key))
		if err != nil {
			return ""
		}
	*/
	//  fmt.Println("delete key",key)
	//  delete(database.OffsetMap,key)
	database.OffsetMap.Delete(key)
	return ""
}

func PopList(key string) string {
	//initCache()
	conn := Pool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("RPOP", key))
	if err != nil {
		//fmt.Println(err)
		return ""
	}
	return value
}
func PushList(key string, value string) {
	//initCache()
	conn := Pool.Get()
	defer conn.Close()
	_, err := conn.Do("Lpush", key, value)
	if err != nil {
		fmt.Println(err)
	}
}
func Ping() error {

	conn := Pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}
