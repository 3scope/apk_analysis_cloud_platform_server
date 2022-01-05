package repository

import (
	"errors"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/util"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// The parameter passed in can be any attribute value of the entity.
type UserRepositoryInterface interface {
	GetTotal(request *Request) (int64, error)
	List(request *Request) ([]model.User, error)
	Add(request *Request) (*model.User, error)
	GetOne(request *Request) (*model.User, error)
	IsExist(request *Request) (int64, error)
	Delete(request *Request) (bool, error)
	Update(request *Request) (*model.User, error)
}

func (repo *UserRepository) GetTotal(request *Request) (int64, error) {
	// The default database.
	var total int64
	db := repo.DB
	var users []model.User
	// The Where object is a struct pointer.
	if request.Entity != nil {
		db = db.Where(request.Entity)
	}
	if err := db.Find(&users).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (repo *UserRepository) List(request *Request) ([]model.User, error) {
	users := make([]model.User, 0)
	db := repo.DB
	// First to get the total number of data.
	total, err := repo.GetTotal(request)
	if err != nil {
		return nil, err
	}
	// Get the correct number of data.
	limit, offset := util.PaginationCheck(request.PageNumber, request.PageSize, int(total), 100)
	if request.Entity != nil {
		db = db.Where(request.Entity)
	}

	if err := db.Order("id asc").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Add(request *Request) (*model.User, error) {
	var user model.User
	// To judge whether the request entity is existed.
	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("user already exists")
	}

	if entity, ok := (request.Entity).(*UserEntity); !ok {
		return nil, errors.New("invalid type entity")
	} else {
		user = (*entity).User
	}

	if err := repo.DB.Model(&user).Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetOne(request *Request) (*model.User, error) {
	// Get only one instance.
	var user model.User
	db := repo.DB

	if request.Entity != nil {
		db = db.Where(request.Entity)
	}

	if err := db.Limit(1).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) IsExist(request *Request) (int64, error) {
	var count int64
	// The "Entity" is pointer.
	if err := repo.DB.Where(request.Entity).Find(&model.User{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (repo *UserRepository) Delete(request *Request) (bool, error) {
	if err := repo.DB.Model(&model.User{}).Where(request.Entity).Delete(request.Entity).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *UserRepository) Update(request *Request) (*model.User, error) {
	var user model.User
	// The verification task is handed over to the front end.
	// Use scan to store the result.
	if err := repo.DB.Model(&model.User{}).Where(request.Entity).Updates(model.User{
		Username:    user.Username,
		RealName:    user.RealName,
		Email:       user.Email,
		Role:        user.Role,
		Description: user.Description,
		Password:    user.Password,
	}).Scan(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
