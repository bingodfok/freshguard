package repertories

import (
	"github.com/bingodfok/freshguard/application/system/internal/model/entity"
	"xorm.io/xorm"
)

type UserQuery struct {
	db *xorm.Engine
}

func NewUserQuery(db *xorm.Engine) *UserQuery {
	return &UserQuery{db: db}
}

func (query *UserQuery) AddUser(user *entity.User) bool {
	count, err := query.db.Insert(user)
	if err != nil {
		return false
	}
	return count > 0
}

func (query *UserQuery) GetUserById(id int64) *entity.User {
	user := &entity.User{}
	has, _ := query.db.Where("id=?", id).Get(user)
	if has {
		return user
	}
	return nil
}

func (query *UserQuery) GetUserByPhone(phone string) *entity.User {
	user := &entity.User{}
	has, _ := query.db.Where("phone=?", phone).Get(user)
	if has {
		return user
	}
	return nil
}

func (query *UserQuery) GetUserDetail(user *entity.User) *entity.User {
	has, _ := query.db.Get(&user)
	if has {
		return user
	}
	return nil
}
