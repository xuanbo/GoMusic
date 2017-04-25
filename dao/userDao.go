package dao

import (
    "github.com/xuanbo/GoMusic/conf"
    "github.com/xuanbo/GoMusic/model"
)

func init()  {
    GetUserDao = UserDaoImpl{}
}

var (
    GetUserDao UserDao
)

type UserDao interface {
    Save(user *model.User)
    Remove(id int64)
}

type UserDaoImpl struct {
}

/*****************************************************
 * implements UserDao interface
 ****************************************************/

func (*UserDaoImpl) Save(user *model.User)  {
    conf.Engine.Insert(user)
}

func (*UserDaoImpl) Remove(id int64) {
    user := new(model.User)
    conf.Engine.Id(id).Delete(user)
}