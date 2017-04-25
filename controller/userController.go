package controller

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/xuanbo/GoMusic/conf"
    "github.com/xuanbo/GoMusic/util"
    "github.com/xuanbo/GoMusic/service"
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
    integer, success := util.ParseToInt64(id)
    if !success {
        return util.NewResult(c, util.ErrorResult(http.StatusBadRequest, "id必须是int", nil))
    }
    user, has := service.GetUserService.Find(integer)
    if !has {
        return util.NewResult(c, util.ErrorResult(http.StatusOK, "找不到user", nil))
    }
    return util.NewResult(c, util.SuccessResult("", user))
}

func remove(c echo.Context) error {
    id := c.Param("id")
    integer, success := util.ParseToInt64(id)
    if !success {
        return util.NewResult(c, util.ErrorResult(http.StatusBadRequest, "id必须是int", nil))
    }
    service.GetUserService.Remove(integer)
    return util.NewResult(c, util.SuccessResult("成功删除user", nil))
}