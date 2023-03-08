package main

import (
	"GinChat/models"
	"GinChat/router"
	"GinChat/utils"
	"github.com/spf13/viper"
	"time"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// 初始化定时器
func InitTimer() {
	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second, models.CleanConnection, "")
}
