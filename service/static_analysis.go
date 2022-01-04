package service

import (
	"errors"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"
)

type StaticAnalysisService struct {
	Repository *repository.StaticAnalysisRepository
}

type StaticAnalysisServiceInterface interface {
	GetTotal(request *repository.Request) (uint, error)
	List(request *repository.Request) ([]model.StaticAnalysis, error)
	Add(request *repository.Request) (*model.StaticAnalysis, error)
	GetOne(request *repository.Request) (*model.StaticAnalysis, error)
	// Someting difference.
	IsExist(request *repository.Request) (bool, error)
	Delete(request *repository.Request) (bool, error)
	Update(request *repository.Request) (*model.StaticAnalysis, error)
}

func (srv *StaticAnalysisService) GetTotal(request *repository.Request) (uint, error) {
	total, err := srv.Repository.GetTotal(request)

	return uint(total), err
}

// To Get Number of instances.
func (srv *StaticAnalysisService) List(request *repository.Request) ([]model.StaticAnalysis, error) {
	return srv.Repository.List(request)
}

// Hand the verification to the front end.
func (srv *StaticAnalysisService) Add(request *repository.Request) (*model.StaticAnalysis, error) {
	return srv.Repository.Add(request)
}

func (srv *StaticAnalysisService) GetOne(request *repository.Request) (*model.StaticAnalysis, error) {
	return srv.Repository.GetOne(request)
}

// To add a middleware.
func (srv *StaticAnalysisService) IsExist(request *repository.Request) (bool, error) {
	count, err := srv.Repository.IsExist(request)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (srv *StaticAnalysisService) Delete(request *repository.Request) (bool, error) {
	if request.Entity == nil {
		return false, errors.New("request parameter error, the 'Entity' attribute is null")
	}
	if staticAnalysis, ok := request.Entity.(model.StaticAnalysis); !ok {
		return false, errors.New("request parameter error, the 'Entity' attribute type error")
	} else if staticAnalysis.ID == 0 {
		return false, errors.New("request parameter error, the primary key is null")
	}
	return srv.Repository.Delete(request)
}

func (srv *StaticAnalysisService) Update(request *repository.Request) (*model.StaticAnalysis, error) {
	result, err := srv.Repository.Update(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
