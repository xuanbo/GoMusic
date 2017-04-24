package main

import (
    "log"
    _ "github.com/go-sql-driver/mysql" // import mysql
    "github.com/xuanbo/GoMusic/conf"
    "github.com/xuanbo/GoMusic/controller"
    "github.com/xuanbo/GoMusic/model"
)

func init() {

}

func main() {
    log.Println("run main...")
    conf.InitDB()
    syncDB()
    conf.InitServer()
    registerController()
	conf.RunServer()
}

// sync db
func syncDB()  {
    log.Println("sync DB.")
    err := conf.Engine.Sync2(new(model.User))
    if err != nil {
        log.Fatalf("failure sync DB，error[%s]\n", err.Error())
    }
}

// register controller
func registerController()  {
    log.Println("register controllers.")
    conf.RegisterControllerRoutes(controller.UserControllerRouter)
}