package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"namilabs/models"
)

// operations for Admin
type AdminController struct {
	baseController
}


func (c *AdminController) Dashboard() {
	c.setupView("admin/pages/dashboard")
  // This page requires login
  if c.isLogged == false {
    c.Redirect("/auth/login", 302)
    return
  }

  // get
  session := c.GetSession("user")
  username := session.(map[string]interface {})
  user := &models.AuthUser{
    Username: username["username"].(string),
  }
  o := orm.NewOrm()
  o.Using("default")
  errRead := o.Read(user, "Username")
  if errRead == orm.ErrNoRows {
    fmt.Println("User not found")
    c.Redirect("/auth/login", 302)
    return
  }
}

