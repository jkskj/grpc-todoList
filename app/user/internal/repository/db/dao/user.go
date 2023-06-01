package dao

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"grpc-todolist/app/user/internal/repository/db/model"
	userPb "grpc-todolist/idl/pb/user"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) CheckUserExist(req *userPb.UserRequest, user *model.User) bool {
	if err := dao.Where("user_name=?", req.UserName).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (dao *UserDao) ShowUserInfo(req *userPb.UserRequest) (user *model.User, err error) {
	//if exist := dao.CheckUserExist(req, user); exist {
	//	return
	//}
	//return nil, errors.New("UserName Not Exist")

	err = dao.Model(&model.User{}).Where("user_name=?", req.UserName).
		First(&user).Error

	return
}

func (dao *UserDao) Create(req *userPb.UserRequest) (user *model.User, err error) {
	var count int64
	dao.Model(&model.User{}).Where("user_name = ?", req.UserName).Count(&count)
	if count != 0 {
		return nil, errors.New("UserName Exist")
	}

	user = &model.User{
		UserName: req.UserName,
		NickName: req.NickName,
	}
	_ = user.SetPassword(req.Password)
	if err = dao.Model(&model.User{}).Create(&user).Error; err != nil {
		fmt.Println("Insert User Error:" + err.Error())
		return nil, err
	}

	return
}
