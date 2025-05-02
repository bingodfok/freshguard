package logic

import (
	"github.com/bingodfok/freshguard/application/system/internal/alc"
	"github.com/bingodfok/freshguard/application/system/internal/common/code"
	"github.com/bingodfok/freshguard/application/system/internal/common/constant"
	"github.com/bingodfok/freshguard/application/system/internal/model/entity"
	"github.com/bingodfok/freshguard/application/system/internal/storage/repertories"
	"github.com/bingodfok/freshguard/pkg/common/errorx"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceLogic struct {
	appCtx *alc.ApplicationContext
}

func NewUserServiceLogic(appCtx *alc.ApplicationContext) *UserServiceLogic {
	return &UserServiceLogic{appCtx: appCtx}
}

// UserPasswordLogin 用户密码登录
func (logic *UserServiceLogic) UserPasswordLogin(username string, password string) (*entity.User, error) {
	query := repertories.NewUserQuery(logic.appCtx.DB)
	user := query.GetUserByPhone(username)
	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errorx.NewCodeError(code.MatchIngError, "用户名或密码错误")
	}
	return user, nil
}

// UserPhoneLogin 手机号登录，手机号没有关联用户则会创建一个
func (logic *UserServiceLogic) UserPhoneLogin(phone string) *entity.User {
	query := repertories.NewUserQuery(logic.appCtx.DB)
	user := query.GetUserByPhone(phone)
	if user == nil {
		user = &entity.User{
			Phone:  phone,
			Name:   "User_" + phone[len(phone)-6:],
			Gender: constant.GenderUnknown,
			Avatar: constant.DefaultAvatar,
		}
		query.AddUser(user)
	}
	return user
}

// QueryUserDetailById 用户ID查询用户详细信息
func (logic *UserServiceLogic) QueryUserDetailById(id int64) (*entity.User, error) {
	query := repertories.NewUserQuery(logic.appCtx.DB)
	user := query.GetUserById(id)
	if user == nil {
		return nil, errorx.NewCodeError(code.NotFoundError, "用户不存在")
	}
	return user, nil
}

// UpdateUserById 更新用户信息
func (logic *UserServiceLogic) UpdateUserById(user *entity.User) (*entity.User, error) {
	query := repertories.NewUserQuery(logic.appCtx.DB)
	userById := query.GetUserById(user.Base.Id)
	if userById == nil {
		return nil, errorx.NewCodeError(code.NotFoundError, "用户不存在")
	}
	if query.UpdateUserById(user) {
		return query.GetUserById(user.Base.Id), nil
	}
	return nil, errorx.NewCodeError(code.UpdateError, "用户信息更新失败")
}

// UpdateUserPassword 修改密码
func (logic *UserServiceLogic) updateUserPassword(id int64, password string) error {
	query := repertories.NewUserQuery(logic.appCtx.DB)
	user := query.GetUserById(id)
	if user == nil {
		return errorx.NewCodeError(code.NotFoundError, "用户不存在")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errorx.NewDefaultCodeError("系统错误")
	}
	user.Password = string(bytes)
	if query.UpdateUserById(user) {
		return nil
	}
	return errorx.NewCodeError(code.UpdateError, "密码更新失败")
}
