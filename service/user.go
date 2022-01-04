package service

import (
	"errors"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"
)

type UserService struct {
	Repository *repository.UserRepository
}

type UserServiceInterface interface {
	GetTotal(request *repository.Request) (uint, error)
	List(request *repository.Request) ([]model.User, error)
	Add(request *repository.Request) (*model.User, error)
	GetOne(request *repository.Request) (*model.User, error)
	// Someting difference.
	IsExist(request *repository.Request) (bool, error)
	Delete(request *repository.Request) (bool, error)
	Update(request *repository.Request) (*model.User, error)
}

func (srv *UserService) GetTotal(request *repository.Request) (uint, error) {
	total, err := srv.Repository.GetTotal(request)

	return uint(total), err
}

// To Get Number of instances.
func (srv *UserService) List(request *repository.Request) ([]model.User, error) {
	return srv.Repository.List(request)
}

// Hand the verification to the front end.
func (srv *UserService) Add(request *repository.Request) (*model.User, error) {
	return srv.Repository.Add(request)
}

func (srv *UserService) GetOne(request *repository.Request) (*model.User, error) {
	return srv.Repository.GetOne(request)
}

// To add a middleware.
func (srv *UserService) IsExist(request *repository.Request) (bool, error) {
	count, err := srv.Repository.IsExist(request)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (srv *UserService) Delete(request *repository.Request) (bool, error) {
	if request.Entity == nil {
		return false, errors.New("request parameter error, the 'Entity' attribute is null")
	}
	if user, ok := request.Entity.(model.User); !ok {
		return false, errors.New("request parameter error, the 'Entity' attribute type error")
	} else if user.ID == 0 {
		return false, errors.New("request parameter error, the primary key is null")
	}
	return srv.Repository.Delete(request)
}

// TODO: Distinguish between administrators and non-administrators.
func (srv *UserService) Update(request *repository.Request) (*model.User, error) {
	result, err := srv.Repository.Update(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
