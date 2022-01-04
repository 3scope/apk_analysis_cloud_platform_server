package repository

import (
	"errors"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/util"
	"gorm.io/gorm"
)

type DynamicAnalysisRepository struct {
	DB *gorm.DB
}

// The parameter passed in can be any attribute value of the entity.
type DynamicAnalysisRepositoryInterface interface {
	GetTotal(request *Request) (int64, error)
	List(request *Request) ([]model.DynamicAnalysis, error)
	Add(request *Request) (*model.DynamicAnalysis, error)
	GetOne(request *Request) (*model.DynamicAnalysis, error)
	IsExist(request *Request) (int64, error)
	Delete(request *Request) (bool, error)
	Update(request *Request) (*model.DynamicAnalysis, error)
}

func (repo *DynamicAnalysisRepository) GetTotal(request *Request) (int64, error) {
	// The default database.
	var total int64
	db := repo.DB
	var dynamicAnalysiss []model.DynamicAnalysis
	// The Where object is a struct pointer.
	if request.Entity != nil {
		db = db.Where(request.Entity)
	}
	if err := db.Find(&dynamicAnalysiss).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (repo *DynamicAnalysisRepository) List(request *Request) ([]model.DynamicAnalysis, error) {
	dynamicAnalysiss := make([]model.DynamicAnalysis, 0)
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

	if err := db.Order("id asc").Limit(limit).Offset(offset).Find(&dynamicAnalysiss).Error; err != nil {
		return nil, err
	}
	return dynamicAnalysiss, nil
}

func (repo *DynamicAnalysisRepository) Add(request *Request) (*model.DynamicAnalysis, error) {
	var dynamicAnalysis model.DynamicAnalysis
	// To judge whether the request entity is existed.
	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("dynamicAnalysis already exists")
	}

	// Use scan to store the result.
	if err := repo.DB.Create(request.Entity).Scan(&dynamicAnalysis).Error; err != nil {
		return nil, err
	}
	return &dynamicAnalysis, nil
}

func (repo *DynamicAnalysisRepository) GetOne(request *Request) (*model.DynamicAnalysis, error) {
	// Get only one instance.
	var dynamicAnalysis model.DynamicAnalysis
	db := repo.DB

	if request.Entity != nil {
		db = db.Where(request.Entity)
	}

	if err := db.Limit(1).Find(&dynamicAnalysis).Error; err != nil {
		return nil, err
	}
	return &dynamicAnalysis, nil
}

func (repo *DynamicAnalysisRepository) IsExist(request *Request) (int64, error) {
	var count int64
	// The "Entity" is pointer.
	if err := repo.DB.Where(request.Entity).Find(&model.DynamicAnalysis{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (repo *DynamicAnalysisRepository) Delete(request *Request) (bool, error) {
	if err := repo.DB.Model(&model.DynamicAnalysis{}).Where(request.Entity).Delete(request.Entity).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *DynamicAnalysisRepository) Update(request *Request) (*model.DynamicAnalysis, error) {
	var dynamicAnalysis model.DynamicAnalysis
	// The verification task is handed over to the front end.
	// Use scan to store the result.
	if err := repo.DB.Model(&model.DynamicAnalysis{}).Where(request.Entity).Updates(model.DynamicAnalysis{
		AnalysisData: dynamicAnalysis.AnalysisData,
		VideoID:      dynamicAnalysis.VideoID,
		CaseID:       dynamicAnalysis.CaseID,
	}).Scan(&dynamicAnalysis).Error; err != nil {
		return nil, err
	}

	return &dynamicAnalysis, nil
}
