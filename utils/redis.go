package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"time"
)

var redisClients *redis.Pool

func CreateTicket() string {
	return TicketName + time.Now().Format("20060102150405000")
}

func RedisSet(key string, val interface{}, ex, seconds string) (reply interface{}, err error) {
	logs.Debug("begin  to RedisSet,key:%s,val:%v,ex:%s,seconds:%s", key, val, ex, seconds)
	conn := getRedisConn()
	defer conn.Close()
	v, err := conn.Do("set", key, val, ex, seconds)
	if err != nil {
		logs.Error("redis.Do err:%v", err)

	}
	logs.Debug("begin  to RedisSet")
	return v, err
}
func RedisDel(key string) (reply interface{}, err error) {
	logs.Debug("begin  to RedisDel,key:%s", key)
	conn := getRedisConn()
	defer conn.Close()
	v, err := conn.Do("del", key)
	if err != nil {
		logs.Error("redis.Do err:%v", err)

	}
	logs.Debug("end to RedisDel")
	return v, err
}

func RedisDo(commandName string, args ...interface{}) (reply interface{}, err error) {
	logs.Debug("begin to redisDo,comm:%s,ars:%v", commandName, args)
	conn := getRedisConn()
	defer conn.Close()
	v, err := conn.Do(commandName, args)
	if err != nil {
		logs.Error("redis.Do err:%v", err)

	}
	logs.Info("end to redisDo")
	return v, err
}

//返回一个redis 连接
func getRedisConn() redis.Conn {
	if redisClients == nil {
		initRedis()
	}
	return redisClients.Get()
}

func initRedis() {
	beego.Info("begin initRedis")
	// 从配置文件获取redis的ip以及db
	REDIS_HOST := beego.AppConfig.DefaultString("redis.host", "127.0.0.1:6379")
	REDIS_DB := beego.AppConfig.DefaultInt("redis.db", 1)
	Passwod := beego.AppConfig.DefaultString("redis.password", "paic1234")
	// 建立连接池

	redisClients = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 80 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", Passwod); err != nil {
				c.Close()
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
	beego.Info("begin initRedis")
}
