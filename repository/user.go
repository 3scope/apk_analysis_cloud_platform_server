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
	// The Where object is a struct pointer.
	// Get the entity pointer.
	user := GetUserInstance(request.Entity)
	if user == nil {
		return 0, errors.New("invalid type of entity")
	}
	// Query conditions can be added here.
	db = db.Debug().Table("users").Where(user)

	if err := db.Count(&total).Error; err != nil {
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

	// Get the entity pointer.
	user := GetUserInstance(request.Entity)
	if user == nil {
		return nil, errors.New("invalid type of entity")
	}
	// The entity is pointer.
	db = db.Debug().Table("users").Where(user)

	if err := db.Order("id ASC").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Add(request *Request) (*model.User, error) {
	// To judge whether the request entity is existed.
	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("user already exists")
	}

	// Get the entity pointer.
	user := GetUserInstance(request.Entity)
	if user == nil {
		return nil, errors.New("invalid type of entity")
	}

	// The "user" variable is pointer.
	if err := repo.DB.Debug().Table("users").Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetOne(request *Request) (*model.User, error) {
	// Get only one instance.
	db := repo.DB
	var count int64

	user := GetUserInstance(request.Entity)
	if user == nil {
		return nil, errors.New("invalid type of entity")
	}
	db = db.Debug().Table("users").Where(user)

	if err := db.Limit(1).Find(user).Count(&count).Error; err != nil {
		return nil, err
	}
	// To judge whether there are result.
	if count == 0 {
		return nil, nil
	}
	return user, nil
}

func (repo *UserRepository) IsExist(request *Request) (int64, error) {
	var count int64

	// Get the entity pointer.
	user := GetUserInstance(request.Entity)
	if user == nil {
		return 0, errors.New("invalid type of entity")
	}

	// The "Entity" is pointer.
	if err := repo.DB.Debug().Table("users").Where(user, "username").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *UserRepository) Delete(request *Request) (bool, error) {
	user := GetUserInstance(request.Entity)
	if user == nil {
		return false, errors.New("invalid type of entity")
	}
	entity, _ := request.Entity.(*UserEntity)
	user.ID = entity.UserID

	if err := repo.DB.Debug().Table("users").Where(user).Delete(user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *UserRepository) Update(request *Request) (*model.User, error) {
	user := GetUserInstance(request.Entity)
	if user == nil {
		return nil, errors.New("invalid type of entity")
	}
	entity, _ := request.Entity.(*UserEntity)
	user.ID = entity.UserID

	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("user already exists")
	}

	// The verification task is handed over to the front end.
	// Cannot change password use this function.
	// Use scan to store the result.
	if err := repo.DB.Debug().Table("users").Where("id = ?", user.ID).Updates(model.User{
		Username:    user.Username,
		RealName:    user.RealName,
		Email:       user.Email,
		Role:        user.Role,
		Description: user.Description,
	}).Scan(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// TODO: Change password function.

// To get a entity instance.
func GetUserInstance(entity interface{}) *model.User {
	var user model.User
	if entity, ok := entity.(*UserEntity); !ok {
		return nil
	} else {
		user = (*entity).User
	}
	return &user
}
