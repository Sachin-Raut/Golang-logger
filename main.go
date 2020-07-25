package main

import (
	"github.com/Sachin-Raut/Golang-logger/log/zap"
)

func main() {

	//always use zap because its fast than logrus

	zap.Info("general info",
		zap.Field("clientID", "10"),
		zap.Field("status", "pending"))
	var err error
	zap.Error("error message", err,
		zap.Field("clientID", "10"),
		zap.Field("status", "pending"))

}
