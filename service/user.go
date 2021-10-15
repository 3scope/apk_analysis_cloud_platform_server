package service

import (
	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"
)

// TODO: Modified to interface-oriented programming.
type UserService struct {
	Repository *repository.UserRepository
}

// To call an existing implementation.
func (srv *UserService) GetTotal(request *repository.Query) (total int) {
	return int(srv.Repository.GetTotal(request))
}

func (srv *UserService) List(
	request *repository.Query) (users []model.User) {
	return srv.Repository.List(request)
}

// Level of docking with the front end.
func (srv *UserService) Add(user model.User) (bool, string) {
	if srv.IsExist(user) {
		return false, "user already exists"
	}
	_, err := srv.Repository.Add(user)
	if err != nil {
		return false, err.Error()
	}
	return true, "ok"
}

func (srv *UserService) Get(user model.User) *model.User {
	return srv.Repository.Get(user)
}

func (srv *UserService) IsExist(user model.User) bool {
	value := srv.Repository.IsExist(user)
	return value != nil
}

func (srv *UserService) Delete(user model.User) (bool, string) {
	if srv.Repository.Delete(user) {
		return true, "ok"
	}
	return false, "delete user failed"
}

func (srv *UserService) Update(user model.User) (bool, string) {
	_, err := srv.Repository.Update(user)
	if err != nil {
		return false, err.Error()
	}
	return true, "ok"
}
