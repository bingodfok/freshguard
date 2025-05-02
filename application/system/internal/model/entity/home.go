package entity

import "github.com/bingodfok/freshguard/pkg/model/entity"

const TableNameHome = "home"

type Home struct {
	Belong     int64
	Name       string
	BaseEntity entity.BaseEntity `xorm:"extends"`
}

func (*Home) TableName() string {
	return TableNameHome
}
