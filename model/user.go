package model

// define User
type User struct {
    Id int64 `xorm:"pk autoincr" json:"id"`
    Name string `xorm:"varchar(50) notnull" json:"name"`
    Hide bool `xorm:"Bool notnull" json:"hide"`
}