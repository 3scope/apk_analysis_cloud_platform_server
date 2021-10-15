package service

import (
	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"
)

type VideoService struct {
	Repository *repository.VideoRepository
}

// To call an existing implementation.
func (srv *VideoService) GetTotal(request *repository.Query) (total int) {
	return int(srv.Repository.GetTotal(request))
}

func (srv *VideoService) List(
	request *repository.Query) (videos []model.VideoStorage) {
	return srv.Repository.List(request)
}

// Level of docking with the front end.
func (srv *VideoService) Add(video model.VideoStorage) (bool, string) {
	if srv.IsExist(video) {
		return false, "video already exists"
	}
	_, err := srv.Repository.Add(video)
	if err != nil {
		return false, err.Error()
	}
	return true, "ok"
}

func (srv *VideoService) Get(video model.VideoStorage) *model.VideoStorage {
	return srv.Repository.Get(video)
}

func (srv *VideoService) IsExist(video model.VideoStorage) bool {
	value := srv.Repository.IsExist(video)
	return value != nil
}

func (srv *VideoService) Delete(video model.VideoStorage) (bool, string) {
	if srv.Repository.Delete(video) {
		return true, "ok"
	}
	return false, "delete video failed"
}

func (srv *VideoService) Update(video model.VideoStorage) (bool, string) {
	_, err := srv.Repository.Update(video)
	if err != nil {
		return false, err.Error()
	}
	return true, "ok"
}
