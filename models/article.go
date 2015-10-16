package models
import (
    "github.com/astaxie/beego/orm"
    "time"
)
type Article struct {
    Id          int
    Title   	string `valid:"Required;MaxSize(25)"`
    Slug    	string `valid:"Required"`
    Summary    	string `orm:"null"`
    Content    	string `orm:"type(text)"`
    User  		*AuthUser  `orm:"rel(fk)"`
    C_date    	time.Time `orm:"auto_now_add;type(datetime)"`
    M_date    	time.Time `orm:"type(datetime)"`
}

func init() {
    orm.RegisterModel(new(Article))
}