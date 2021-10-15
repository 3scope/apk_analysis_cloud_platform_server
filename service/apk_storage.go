package service

import (
	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"
)

type ApkService struct {
	Repository *repository.ApkRepository
}

// To call an existing implementation.
func (srv *ApkService) GetTotal(request *repository.Query) (total int) {
	return int(srv.Repository.GetTotal(request))
}

func (srv *ApkService) List(
	request *repository.Query) (apks []model.ApkStorage) {
	return srv.Repository.List(request)
}

// Level of docking with the front end.
func (srv *ApkService) Add(apk model.ApkStorage) (bool, string) {
	if srv.IsExist(apk) {
		return false, "apk already exists"
	}
	_, err := srv.Repository.Add(apk)
	if err != nil {
		return false, err.Error()
	}
	return true, "ok"
}

func (srv *ApkService) Get(apk model.ApkStorage) *model.ApkStorage {
	return srv.Repository.Get(apk)
}

func (srv *ApkService) IsExist(apk model.ApkStorage) bool {
	value := srv.Repository.IsExist(apk)
	return value != nil
}

func (srv *ApkService) Delete(apk model.ApkStorage) (bool, string) {
	if srv.Repository.Delete(apk) {
		return true, "ok"
	}
	return false, "delete apk failed"
}

func (srv *ApkService) Update(apk model.ApkStorage) (bool, string) {
	_, err := srv.Repository.Update(apk)
	if err != nil {
		return false, err.Error()
	}
	return true, "ok"
}
