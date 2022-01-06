package service

import (
	"errors"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"
)

type DynamicAnalysisService struct {
	Repository *repository.DynamicAnalysisRepository
}

type DynamicAnalysisServiceInterface interface {
	GetTotal(request *repository.Request) (uint, error)
	List(request *repository.Request) ([]model.DynamicAnalysis, error)
	Add(request *repository.Request) (*model.DynamicAnalysis, error)
	GetOne(request *repository.Request) (*model.DynamicAnalysis, error)
	// Someting difference.
	IsExist(request *repository.Request) (bool, error)
	Delete(request *repository.Request) (bool, error)
	Update(request *repository.Request) (*model.DynamicAnalysis, error)
}

func (srv *DynamicAnalysisService) GetTotal(request *repository.Request) (uint, error) {
	total, err := srv.Repository.GetTotal(request)

	return uint(total), err
}

// To Get Number of instances.
func (srv *DynamicAnalysisService) List(request *repository.Request) ([]model.DynamicAnalysis, error) {
	return srv.Repository.List(request)
}

// Hand the verification to the front end.
func (srv *DynamicAnalysisService) Add(request *repository.Request) (*model.DynamicAnalysis, error) {
	return srv.Repository.Add(request)
}

func (srv *DynamicAnalysisService) GetOne(request *repository.Request) (*model.DynamicAnalysis, error) {
	return srv.Repository.GetOne(request)
}

// To add a middleware.
func (srv *DynamicAnalysisService) IsExist(request *repository.Request) (bool, error) {
	count, err := srv.Repository.IsExist(request)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (srv *DynamicAnalysisService) Delete(request *repository.Request) (bool, error) {
	if request.Entity == nil {
		return false, errors.New("request parameter error, the 'Entity' attribute is null")
	}
	// Determine whether it is a pointer.
	if dynamicAnalysis, ok := request.Entity.(*(repository.DynamicAnalysisEntity)); !ok {
		return false, errors.New("request parameter error, the 'Entity' attribute type error")
	} else if dynamicAnalysis == nil {
		return false, errors.New("request parameter error, the entity is null")
	}
	return srv.Repository.Delete(request)
}

func (srv *DynamicAnalysisService) Update(request *repository.Request) (*model.DynamicAnalysis, error) {
	result, err := srv.Repository.Update(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
