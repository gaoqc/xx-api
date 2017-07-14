package models

import (
	"time"
)

type Model struct {
	Id int `orm:"pk;auto"`
}
type TimeModel struct {
	Model
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
