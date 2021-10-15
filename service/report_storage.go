package service

import (
	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"
)

type ReportService struct {
	Repository *repository.ReportRepository
}

// To call an existing implementation.
func (srv *ReportService) GetTotal(request *repository.Query) (total int) {
	return int(srv.Repository.GetTotal(request))
}

func (srv *ReportService) List(
	request *repository.Query) (reports []model.ReportStorage) {
	return srv.Repository.List(request)
}

// Level of docking with the front end.
func (srv *ReportService) Add(report model.ReportStorage) (bool, string) {
	if srv.IsExist(report) {
		return false, "report already exists"
	}
	_, err := srv.Repository.Add(report)
	if err != nil {
		return false, err.Error()
	}
	return true, "ok"
}

func (srv *ReportService) Get(report model.ReportStorage) *model.ReportStorage {
	return srv.Repository.Get(report)
}

func (srv *ReportService) IsExist(report model.ReportStorage) bool {
	value := srv.Repository.IsExist(report)
	return value != nil
}

func (srv *ReportService) Delete(report model.ReportStorage) (bool, string) {
	if srv.Repository.Delete(report) {
		return true, "ok"
	}
	return false, "delete report failed"
}

func (srv *ReportService) Update(report model.ReportStorage) (bool, string) {
	_, err := srv.Repository.Update(report)
	if err != nil {
		return false, err.Error()
	}
	return true, "ok"
}
