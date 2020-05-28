package pkg

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var (
	cfg *ini.File
	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
)

func init(){
	var err error
	var dataFile = "/users/xiashuang/go/src/awesomeProject/src/config/app.ini"
	cfg,err = ini.Load(dataFile)
	if err != nil {
		logrus.Info(err)
	}
	RunMode = cfg.Section("").Key("run_model").MustString("debug")
	loadServer()
	loadApp()
}

func loadServer(){
	server,_ := cfg.GetSection("server")
	HTTPPort = server.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout =  time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadApp(){
	_, err := cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
}
