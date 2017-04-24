package conf

import (
    "log"
    "net/http"
    "github.com/labstack/echo"
)

// define Route
type Route struct {
    Patten string // path patten
    Method string // http method such as GET、POST、PUT、DELETE
    Handler func(c echo.Context) error // `echo framework` required func type
}

// define Controller Route
type ControllerRouter struct {
    Name string //  controller name
    Routes []Route // routes of controller
}

// register Controller Route
func RegisterControllerRoutes(controllerRouters ...ControllerRouter) {
    controllerRoutersLength := len(controllerRouters)
    for i := 0; i < controllerRoutersLength; i++ {
        log.Printf("register controller[%s].\n", controllerRouters[i].Name)
        routesLength := len(controllerRouters[i].Routes)
        for j := 0; j < routesLength; j++ {
            registerRoute(&controllerRouters[i].Routes[j])
        }
    }
}

// register Route to `echo framework`
func registerRoute(r *Route) {
    patten := r.Patten
    method := r.Method
    handler := r.Handler
    log.Printf("register route patten[%s]，method[%s].\n",  patten, method)
    // register
    switch method {
        case http.MethodGet:
            Echo.GET(patten, handler)
        case http.MethodPost:
            Echo.POST(patten, handler)
        case http.MethodPut:
            Echo.PUT(patten, handler)
        case http.MethodDelete:
            Echo.DELETE(patten, handler)
        // ignore other http method，such as OPTIONS、PATCH
        default:
            log.Printf("http method[%s] not support.\n", method)
    }
}