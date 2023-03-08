package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	Red *redis.Client
	DB  *gorm.DB
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper err : ", err)
	}

}

func InitMySQL() {
	//自定义日志模版打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL的阈值
			LogLevel:      logger.Info, //日志级别
			Colorful:      true,        //颜色
		},
	)

	DB, _ = gorm.Open(mysql.Open(viper.GetString("MySQL.dsn")), &gorm.Config{Logger: newLogger})

}
func InitRedis() {

	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("Redis.addr"),
		Password:     viper.GetString("Redis.password"),
		DB:           viper.GetInt("Redis.DB"),
		PoolSize:     viper.GetInt("Redis.poolSize"),
		MinIdleConns: viper.GetInt("Redis.minIDleConn"),
	})

}

const PublishKey = "websocket"

// Publish 发布redis消息
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("redis publish msg : ", msg)
	err = Red.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	fmt.Println("redis subscribe msg --ctx : ", ctx)
	msg, err := sub.ReceiveMessage(ctx)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("redis subscribe msg : ", msg.Payload)
	return msg.Payload, err
}
