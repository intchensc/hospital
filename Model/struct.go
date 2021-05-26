package Model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type Patient struct {
	gorm.Model
	Name string `form:"xingming"`
	Sex string `form:"xingbie"`
	Age int `form:"nianling"`
	Tel string `form:"shouji"`
	Disease string `form:"zhenduan"`
}

type Bed struct {
	gorm.Model
	Name string
	Tel string
	Disease string
}
