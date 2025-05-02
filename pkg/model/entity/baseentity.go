package entity

import "time"

type BaseEntity struct {
	Id       int64     `xorm:"pk"`
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	DeleteAt time.Time `xorm:"deleted"`
	Version  int       `xorm:"version"`
}
