package util

import (
    "net/http"
    "github.com/labstack/echo"
)

type Result struct {
    Code int `json:"code"`
    Message string `json:"message"`
    Data interface{} `json:"data"`
}

func SuccessResult(message string, data interface{}) *Result {
    return &Result{http.StatusOK, message, data}
}

func ErrorResult(code int, message string, data interface{}) *Result {
    return &Result{code, message, data}
}

func NewResult(c echo.Context, r *Result) error  {
    return c.JSON(http.StatusOK, r)
}