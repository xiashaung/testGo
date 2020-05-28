package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

func testGo()  {
	go func() {
		for  {
			time.Sleep(1*time.Second)
			logrus.Info("测试协程函数")
		}
	}()
}
