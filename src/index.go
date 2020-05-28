package main

import (
   "fmt"
   "github.com/fvbock/endless"
   "github.com/my/repo/pkg"
   "github.com/my/repo/route"
   "github.com/sirupsen/logrus"
   "syscall"
   "time"
)

func main() {
   startServer();
}


func startServer(){
   endless.DefaultReadTimeOut = 60 * time.Second
   endless.DefaultWriteTimeOut = 60 * time.Second
   endless.DefaultMaxHeaderBytes = 1 << 20
   serverPort := fmt.Sprintf("0.0.0.0:%d", pkg.HTTPPort);
   server :=endless.NewServer(serverPort,route.InitRouter())
   server.BeforeBegin = func (add string){
      logrus.Printf("pid 是: %d",syscall.Getpid())
   }
   _ = server.ListenAndServe()
}


