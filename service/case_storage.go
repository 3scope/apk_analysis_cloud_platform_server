package service

import (
	"errors"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/repository"
)

type CaseService struct {
	Repository *repository.CaseRepository
}

type CaseServiceInterface interface {
	GetTotal(request *repository.Request) (uint, error)
	List(request *repository.Request) ([]model.Case, error)
	Add(request *repository.Request) (*model.Case, error)
	GetOne(request *repository.Request) (*model.Case, error)
	// Someting difference.
	IsExist(request *repository.Request) (bool, error)
	Delete(request *repository.Request) (bool, error)
	Update(request *repository.Request) (*model.Case, error)
}

func (srv *CaseService) GetTotal(request *repository.Request) (uint, error) {
	total, err := srv.Repository.GetTotal(request)

	return uint(total), err
}

// To Get Number of instances.
func (srv *CaseService) List(request *repository.Request) ([]model.Case, error) {
	return srv.Repository.List(request)
}

// Hand the verification to the front end.
func (srv *CaseService) Add(request *repository.Request) (*model.Case, error) {
	return srv.Repository.Add(request)
}

func (srv *CaseService) GetOne(request *repository.Request) (*model.Case, error) {
	return srv.Repository.GetOne(request)
}

// To add a middleware.
func (srv *CaseService) IsExist(request *repository.Request) (bool, error) {
	count, err := srv.Repository.IsExist(request)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (srv *CaseService) Delete(request *repository.Request) (bool, error) {
	if request.Entity == nil {
		return false, errors.New("request parameter error, the 'Entity' attribute is null")
	}
	// Determine whether it is a pointer.
	if caseInstance, ok := request.Entity.(*(repository.CaseEntity)); !ok {
		return false, errors.New("request parameter error, the 'Entity' attribute type error")
	} else if caseInstance == nil {
		return false, errors.New("request parameter error, the entity is null")
	}
	return srv.Repository.Delete(request)
}

func (srv *CaseService) Update(request *repository.Request) (*model.Case, error) {
	result, err := srv.Repository.Update(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
