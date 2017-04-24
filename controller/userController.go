package controller

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/xuanbo/GoMusic/conf"
    "github.com/xuanbo/GoMusic/dao"
    "github.com/xuanbo/GoMusic/model"
)

/*****************************************************
 * define UserController Route info
 ****************************************************/
var UserControllerRouter = conf.ControllerRouter {
    Name: "userController",
    Routes: []conf.Route {
        {
            Patten: "/users",
            Method: http.MethodGet,
            Handler: users,
        },
        {
            Patten: "/user/:id",
            Method: http.MethodGet,
            Handler: show,
        },
        {
            Patten: "/user/:id",
            Method: http.MethodDelete,
            Handler: remove,
        },
    },
}

/*****************************************************
 * define func
 ****************************************************/

func users(c echo.Context) error {
    return c.JSON(http.StatusOK, []int{1, 2, 3})
}

func show(c echo.Context) error {
    id := c.Param("id")
    return c.JSON(http.StatusOK, id)
}

func remove(c echo.Context) error {
    return c.JSON(http.StatusOK, 1)
}