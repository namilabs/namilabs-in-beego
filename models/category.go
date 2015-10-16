package models
import (
    "github.com/astaxie/beego/orm"
    "time"
)
type Category struct {
    Id          int
    Title   	string `valid:"Required;MaxSize(25)"`
    Slug    	string `valid:"Required"`
    Status    	int `valid:"Required"`
    C_date    	time.Time `orm:"auto_now_add;type(datetime)"`
    M_date    	time.Time `orm:"type(datetime);null"`
}

func init() {
    orm.RegisterModel(new(Category))
}