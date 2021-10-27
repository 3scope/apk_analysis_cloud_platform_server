package repository

import (
	"errors"
	"log"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/util"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// The parameter passed in can be any attribute value of the user object.
type UserRepositoryInterface interface {
	GetTotal(request *Query) (total int64)
	List(request *Query) (users []model.User)
	Add(user model.User) (*model.User, error)
	Get(user model.User) *model.User
	IsExist(user model.User) *model.User
	Delete(user model.User) bool
	Update(user model.User) (*model.User, error)
}

func (ur *UserRepository) GetTotal(request *Query) (total int64) {
	db := ur.DB
	var users []model.User
	//TODO: Verify whether where needs to be processed.
	if request.Where != "" {
		db = db.Where(request.Where)
	}
	if err := db.Find(&users).Count(&total).Error; err != nil {
		log.Println(err)
		return 0
	}
	return total
}

func (ur *UserRepository) List(request *Query) (users []model.User) {
	db := ur.DB
	// First to get the total number of data.
	total := ur.GetTotal(request)
	// Get the correct number of data.
	limit, offset := util.PaginationCheck(request.PageNumber, request.PageSize, int(total), 100)
	if request.Where != "" {
		db = db.Where(request.Where)
	}
	if err := db.Order("id asc").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		log.Println(err)
		return nil
	}
	return users
}

func (ur *UserRepository) Add(user model.User) (*model.User, error) {
	if err := ur.DB.Create(&user).Error; err != nil {
		return nil, errors.New("user registration failed")
	}
	return &user, nil
}

func (ur *UserRepository) Get(user model.User) *model.User {
	// It can be one or multiple.
	if err := ur.DB.Where(&user).Find(&user).Error; err != nil {
		log.Println(err)
		return nil
	}
	return &user
}

func (ur *UserRepository) IsExist(user model.User) *model.User {
	if err := ur.DB.Where(&user).Find(&user); err != nil {
		log.Println(err)
		return nil
	}
	return &user
}

func (ur *UserRepository) Delete(user model.User) bool {
	if err := ur.DB.Delete(&user); err != nil {
		log.Println(err)
		return false
	}
	return true
}

// This method can not change the password.
func (ur *UserRepository) Update(user model.User) (*model.User, error) {
	if err := ur.DB.Model(&user).Updates(model.User{Username: user.Username,
		RealName:    user.RealName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Description: user.Description}).Error; err != nil {
		return nil, errors.New("user update failed")
	}
	return &user, nil
}
