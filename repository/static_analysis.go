package repository

import (
	"errors"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/util"
	"gorm.io/gorm"
)

type StaticAnalysisRepository struct {
	DB *gorm.DB
}

// The parameter passed in can be any attribute value of the entity.
type StaticAnalysisRepositoryInterface interface {
	GetTotal(request *Request) (int64, error)
	List(request *Request) ([]model.StaticAnalysis, error)
	Add(request *Request) (*model.StaticAnalysis, error)
	GetOne(request *Request) (*model.StaticAnalysis, error)
	IsExist(request *Request) (int64, error)
	Delete(request *Request) (bool, error)
	Update(request *Request) (*model.StaticAnalysis, error)
}

func (repo *StaticAnalysisRepository) GetTotal(request *Request) (int64, error) {
	// The default database.
	var total int64
	db := repo.DB
	var staticAnalysiss []model.StaticAnalysis
	// The Where object is a struct pointer.
	if request.Entity != nil {
		db = db.Where(request.Entity)
	}
	if err := db.Find(&staticAnalysiss).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (repo *StaticAnalysisRepository) List(request *Request) ([]model.StaticAnalysis, error) {
	staticAnalysiss := make([]model.StaticAnalysis, 0)
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

	if err := db.Order("id asc").Limit(limit).Offset(offset).Find(&staticAnalysiss).Error; err != nil {
		return nil, err
	}
	return staticAnalysiss, nil
}

func (repo *StaticAnalysisRepository) Add(request *Request) (*model.StaticAnalysis, error) {
	var staticAnalysis model.StaticAnalysis
	// To judge whether the request entity is existed.
	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("staticAnalysis already exists")
	}

	// Use scan to store the result.
	if err := repo.DB.Create(request.Entity).Scan(&staticAnalysis).Error; err != nil {
		return nil, err
	}
	return &staticAnalysis, nil
}

func (repo *StaticAnalysisRepository) GetOne(request *Request) (*model.StaticAnalysis, error) {
	// Get only one instance.
	var staticAnalysis model.StaticAnalysis
	db := repo.DB

	if request.Entity != nil {
		db = db.Where(request.Entity)
	}

	if err := db.Limit(1).Find(&staticAnalysis).Error; err != nil {
		return nil, err
	}
	return &staticAnalysis, nil
}

func (repo *StaticAnalysisRepository) IsExist(request *Request) (int64, error) {
	var count int64
	// The "Entity" is pointer.
	if err := repo.DB.Where(request.Entity).Find(&model.StaticAnalysis{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (repo *StaticAnalysisRepository) Delete(request *Request) (bool, error) {
	if err := repo.DB.Model(&model.StaticAnalysis{}).Where(request.Entity).Delete(request.Entity).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *StaticAnalysisRepository) Update(request *Request) (*model.StaticAnalysis, error) {
	var staticAnalysis model.StaticAnalysis
	// The verification task is handed over to the front end.
	// Use scan to store the result.
	if err := repo.DB.Model(&model.StaticAnalysis{}).Where(request.Entity).Updates(model.StaticAnalysis{
		AppName:     staticAnalysis.AppName,
		AppVersion:  staticAnalysis.AppVersion,
		ApkSize:     staticAnalysis.ApkSize,
		PackageName: staticAnalysis.PackageName,
		MD5:         staticAnalysis.MD5,
		SHA160:      staticAnalysis.SHA160,
		SHA256:      staticAnalysis.SHA256,
		SDK:         staticAnalysis.SDK,
		Certificate: staticAnalysis.Certificate,
		Email:       staticAnalysis.Email,
		IPAdress:    staticAnalysis.IPAdress,
		DomainName:  staticAnalysis.DomainName,
		Permission:  staticAnalysis.Permission,
	}).Scan(&staticAnalysis).Error; err != nil {
		return nil, err
	}

	return &staticAnalysis, nil
}
