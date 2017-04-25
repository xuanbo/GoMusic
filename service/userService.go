package service

import (
    "github.com/xuanbo/GoMusic/model"
    "github.com/xuanbo/GoMusic/dao"
)

func init()  {
    GetUserService = &UserServiceImpl{}
}

var (
    GetUserService UserService
)

// define UserService
type UserService interface {
    Save(user *model.User)
    Modify(user *model.User)
    Remove(id int64)
    Find(id int64) (*model.User, bool)
}

// define UserServiceImpl implements UserService
type UserServiceImpl struct {
}

/*****************************************************
 * implements UserService interface
 ****************************************************/

func (*UserServiceImpl) Save(user *model.User) {
    dao.GetUserDao.Save(user)
}

func (*UserServiceImpl) Modify(user *model.User) {
    dao.GetUserDao.Modify(user)
}


func (this *UserServiceImpl) Remove(id int64) {
    user, has := this.Find(id)
    if has {
        user.Hide = true
        this.Modify(user)
    }
}

func (*UserServiceImpl) Find(id int64) (*model.User, bool) {
    return dao.GetUserDao.Find(id)
}