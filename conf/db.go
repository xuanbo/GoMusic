package conf

import (
    "log"
    "github.com/go-xorm/xorm"
)

type DataSource struct {
    DriverName string
    Username string
    Password string
    DataBaseName string
    Url string
}

var (
    dataSource *DataSource
    Engine *xorm.Engine
)

// init DataSource
func initDataSource() {
    dataSource = &DataSource {
        DriverName: "mysql",
        Username: "root",
        Password: "123456",
        DataBaseName: "music",
        Url: "root:123456@tcp(127.0.0.1:3306)/music?charset=utf8",
    }
}

// init xorm engine
func initEngine() {
    var err error
    Engine, err = xorm.NewEngine(dataSource.DriverName, dataSource.Url)
    if err != nil {
        log.Fatalf("failure init xorm engine[%s]\n", err.Error())
    }
}

// init db
func InitDB()  {
    initDataSource()
    initEngine()
}