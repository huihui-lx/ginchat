package main

import (
	"ginchat/models"
	"ginchat/router" 
	"ginchat/utils"
	"time"

	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	InitTimer()
	r := router.Router()                  
	r.Run(viper.GetString("port.server")) 
}

//初始化定时器
func InitTimer() {
	delay := time.Duration(viper.GetInt("timeout.DelayHeartbeat")) * time.Second
	tick := time.Duration(viper.GetInt("timeout.HeartbeatHz")) * time.Second
	utils.Timer(delay, tick, models.CleanConnection, "")
}
