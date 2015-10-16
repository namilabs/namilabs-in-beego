package main

import (
  "fmt"
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql"

  // routers
  _ "namilabs/routers"

  // models imports
  _ "namilabs/models"
)

func init() {
  orm.RegisterDriver("mysql", orm.DR_MySQL)
  orm.RegisterDataBase("default", "mysql", "root:root@/nami_go?charset=utf8")
  name := "default"
  force := false
  verbose := false
  err := orm.RunSyncdb(name, force, verbose)
  if err != nil {
      fmt.Println(err)
  }
}

func main() {
	beego.Run()
}

