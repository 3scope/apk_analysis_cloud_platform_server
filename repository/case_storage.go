package repository

import (
	"errors"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/util"
	"gorm.io/gorm"
)

type CaseRepository struct {
	DB *gorm.DB
}

// The parameter passed in can be any attribute value of the entity.
type CaseRepositoryInterface interface {
	GetTotal(request *Request) (int64, error)
	List(request *Request) ([]model.Case, error)
	Add(request *Request) (*model.Case, error)
	GetOne(request *Request) (*model.Case, error)
	IsExist(request *Request) (int64, error)
	Delete(request *Request) (bool, error)
	Update(request *Request) (*model.Case, error)
}

func (repo *CaseRepository) GetTotal(request *Request) (int64, error) {
	// The default database.
	var total int64
	db := repo.DB
	var caseInstances []model.Case
	// The Where object is a struct pointer.
	if request.Entity != nil {
		db = db.Where(request.Entity)
	}
	if err := db.Find(&caseInstances).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (repo *CaseRepository) List(request *Request) ([]model.Case, error) {
	caseInstances := make([]model.Case, 0)
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

	if err := db.Order("id asc").Limit(limit).Offset(offset).Find(&caseInstances).Error; err != nil {
		return nil, err
	}
	return caseInstances, nil
}

func (repo *CaseRepository) Add(request *Request) (*model.Case, error) {
	var caseInstance model.Case
	// To judge whether the request entity is existed.
	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("caseInstance already exists")
	}

	// Use scan to store the result.
	if err := repo.DB.Create(request.Entity).Scan(&caseInstance).Error; err != nil {
		return nil, err
	}
	return &caseInstance, nil
}

func (repo *CaseRepository) GetOne(request *Request) (*model.Case, error) {
	// Get only one instance.
	var caseInstance model.Case
	db := repo.DB

	if request.Entity != nil {
		db = db.Where(request.Entity)
	}

	if err := db.Limit(1).Find(&caseInstance).Error; err != nil {
		return nil, err
	}
	return &caseInstance, nil
}

func (repo *CaseRepository) IsExist(request *Request) (int64, error) {
	var count int64
	// The "Entity" is pointer.
	if err := repo.DB.Where(request.Entity).Find(&model.Case{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (repo *CaseRepository) Delete(request *Request) (bool, error) {
	if err := repo.DB.Model(&model.Case{}).Where(request.Entity).Delete(request.Entity).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *CaseRepository) Update(request *Request) (*model.Case, error) {
	var caseInstance model.Case
	// The verification task is handed over to the front end.
	// Use scan to store the result.
	if err := repo.DB.Model(&model.Case{}).Where(request.Entity).Updates(model.Case{
		CaseName:          caseInstance.CaseName,
		CaseDate:          caseInstance.CaseDate,
		StaticAnalysisID:  caseInstance.StaticAnalysisID,
		DynamicAnalysisID: caseInstance.DynamicAnalysisID,
		ReportID:          caseInstance.ReportID,
		UploaderID:        caseInstance.UploaderID,
		UploaderUserName:  caseInstance.UploaderUserName,
	}).Scan(&caseInstance).Error; err != nil {
		return nil, err
	}

	return &caseInstance, nil
}
