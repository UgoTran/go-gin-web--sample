package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Engine = gin.Default()
var ServerPort = "SERVER_PORT"

func main() {
	fmt.Println("Setup gin framework")
	port := fmt.Sprintf("%v", viper.Get(ServerPort))
	err := Engine.Run(port)
	logrus.Print("Gin server stared on port " + port)
	if err != nil {
		logrus.Fatal("[Error] failed to start Gin server due to: " + err.Error())
		return
	}
}

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
