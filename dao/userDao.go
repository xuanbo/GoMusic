package dao

import (
    "github.com/xuanbo/GoMusic/conf"
    "github.com/xuanbo/GoMusic/model"
    "github.com/labstack/gommon/log"
)

func init()  {
    GetUserDao = &UserDaoImpl{}
}

var (
    GetUserDao UserDao
)

// define UserDao
type UserDao interface {
    Save(user *model.User)
    Modify(user *model.User)
    Find(id int64) (*model.User, bool)
}

// define UserDaoImpl implements UserDao
type UserDaoImpl struct {
}

/*****************************************************
 * implements UserDao interface
 ****************************************************/

func (*UserDaoImpl) Save(user *model.User) {
    _, err := conf.Engine.Insert(user)
    if err != nil {
        log.Printf("UserDao save DB error[%s]\n", err.Error())
    }
}

func (*UserDaoImpl) Modify(user *model.User) {
    _, err := conf.Engine.Update(user)
    if err != nil {
        log.Printf("UserDao modify DB error[%s]\n", err.Error())
    }
}

func (*UserDaoImpl) Find(id int64) (*model.User, bool)  {
    user := new(model.User)
    has, err := conf.Engine.Id(id).Get(user)
    if err != nil {
        log.Printf("UserDao find DB error[%s]\n", err.Error())
    }
    return user, has
}