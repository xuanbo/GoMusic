package conf

import (
    "time"
    "net/http"
    "github.com/labstack/echo"
)

var (
    server *http.Server
    Echo *echo.Echo
)

func RunServer()  {
    Echo.Logger.Fatal(Echo.StartServer(server))
}

// init server info
func InitServer()  {
    server = &http.Server{
        Addr: ":8080",
        ReadTimeout: 20 * time.Second,
        WriteTimeout: 20 * time.Second,
    }
    initEcho()
}

// init echo framework
func initEcho() {
    Echo = echo.New()
}