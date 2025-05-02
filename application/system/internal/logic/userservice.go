package logic

import (
	"github.com/bingodfok/freshguard/application/system/internal/alc"
	"github.com/bingodfok/freshguard/application/system/internal/model/entity"
	"github.com/bingodfok/freshguard/application/system/internal/storage/repertories"
	"github.com/bingodfok/freshguard/pkg/common/errorx"
	"golang.org/x/net/context"
)

type UserServiceLogic struct {
	appCtx *alc.ApplicationContext
}

func NewUserServiceLogic(appCtx *alc.ApplicationContext) *UserServiceLogic {
	return &UserServiceLogic{appCtx: appCtx}
}

func (logic *UserServiceLogic) UserPhoneLogin(context context.Context, phone string, password string) (*entity.User, error) {
	query := repertories.NewUserQuery(logic.appCtx.DB)
	user := query.GetUserByPhone(phone)
	if user == nil || user.Password != password {
		return nil, errorx.NewCodeError(401, "用户名或密码错误")
	}
	return user, nil
}

func (logic *UserServiceLogic) QueryUserDetailById(context context.Context, id int64) (*entity.User, error) {
	query := repertories.NewUserQuery(logic.appCtx.DB)
	user := query.GetUserById(id)
	if user == nil {
		return nil, errorx.NewDefaultCodeError("没有查询到用户信息")
	}
	return user, nil
}
