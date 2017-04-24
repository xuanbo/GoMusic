package model

// define User
type User struct {
    Id int64 `xorm:"pk autoincr"`
    Name string `xorm:"varchar(50) notnull"`
}