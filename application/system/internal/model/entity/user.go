package entity

import "github.com/bingodfok/freshguard/pkg/model/entity"

const TableNameUser = "user"

type User struct {
	Name     string
	Password string
	Phone    string
	Gender   string
	Avatar   string
	Base     entity.BaseEntity `xorm:"extends"`
}

func (*User) TableName() string {
	return TableNameUser
}
