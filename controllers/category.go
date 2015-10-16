package controllers

import (
  "fmt"
  "strconv"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/validation"
  "github.com/astaxie/beego/orm"
  "namilabs/models"
  "github.com/mrvdot/golang-utils"
  //"github.com/astaxie/beego/utils/pagination"
)

// operations for Admin
type CategoryController struct {
  baseController
}


func (c *CategoryController) Index() {
  c.setupView("admin/category/index")
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


 /* perPage := 20
  paginator := pagination.SetPaginator(c, perPage, CountCategory)

    // fetch the next 20 posts
  this.Data["categories"] = ListCategoryByOffsetAndLimit(paginator.Offset(), perPage)*/
}

func (c *CategoryController) GetAdd() {
  c.setupView("admin/category/form")
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

func (c *CategoryController) PostAdd() {
  c.setupView("admin/category/form")

  if c.Ctx.Input.Method() == "POST" {
    flash := beego.NewFlash()
    valid := validation.Validation{}

    title := c.GetString("title")
    slug := utils.GenerateSlug(title)
    status := c.GetString("status")
    coba, err := strconv.Atoi(status)

    o := orm.NewOrm()
    o.Using("default")
    category := &models.Category{
      Title: title,
      Slug: slug,
      Status: coba,
    }

    b, err := valid.Valid(category)
    if err != nil {
        fmt.Println(err)
    }
    if !b {
      errormap := []string{}
      for _, err := range valid.Errors {
        errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
      }
      flash.Error("Invalid data!")
      flash.Store(&c.Controller)
      c.Data["Errors"] = errormap
      fmt.Println(errormap)
      return
    }

    
    //category := models.Category{}

    id, err := o.Insert(category)
    if err == nil {
        fmt.Println(id)
        flash.Notice("Category has been saved")
        flash.Store(&c.Controller)
        c.Redirect("/nladmin/category", 302)
        return;
    }else{
      msg := fmt.Sprintf("Couldn't insert new article. Reason: ", err)
      beego.Debug(msg)
      flash.Error(msg)
      flash.Store(&c.Controller)
    }

    

  }
}

