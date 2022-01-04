package repository

import (
	"errors"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/util"
	"gorm.io/gorm"
)

type ReportRepository struct {
	DB *gorm.DB
}

// The parameter passed in can be any attribute value of the entity.
type ReportRepositoryInterface interface {
	GetTotal(request *Request) (int64, error)
	List(request *Request) ([]model.Report, error)
	Add(request *Request) (*model.Report, error)
	GetOne(request *Request) (*model.Report, error)
	IsExist(request *Request) (int64, error)
	Delete(request *Request) (bool, error)
	Update(request *Request) (*model.Report, error)
}

func (repo *ReportRepository) GetTotal(request *Request) (int64, error) {
	// The default database.
	var total int64
	db := repo.DB
	var reports []model.Report
	// The Where object is a struct pointer.
	if request.Entity != nil {
		db = db.Where(request.Entity)
	}
	if err := db.Find(&reports).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (repo *ReportRepository) List(request *Request) ([]model.Report, error) {
	reports := make([]model.Report, 0)
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

	if err := db.Order("id asc").Limit(limit).Offset(offset).Find(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}

func (repo *ReportRepository) Add(request *Request) (*model.Report, error) {
	var report model.Report
	// To judge whether the request entity is existed.
	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("report already exists")
	}

	// Use scan to store the result.
	if err := repo.DB.Create(request.Entity).Scan(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (repo *ReportRepository) GetOne(request *Request) (*model.Report, error) {
	// Get only one instance.
	var report model.Report
	db := repo.DB

	if request.Entity != nil {
		db = db.Where(request.Entity)
	}

	if err := db.Limit(1).Find(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (repo *ReportRepository) IsExist(request *Request) (int64, error) {
	var count int64
	// The "Entity" is pointer.
	if err := repo.DB.Where(request.Entity).Find(&model.Report{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (repo *ReportRepository) Delete(request *Request) (bool, error) {
	if err := repo.DB.Model(&model.Report{}).Where(request.Entity).Delete(request.Entity).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *ReportRepository) Update(request *Request) (*model.Report, error) {
	var report model.Report
	// The verification task is handed over to the front end.
	// Use scan to store the result.
	if err := repo.DB.Model(&model.Report{}).Where(request.Entity).Updates(model.Report{
		ReportName: report.ReportName,
		CaseID:     report.CaseID,
	}).Scan(&report).Error; err != nil {
		return nil, err
	}

	return &report, nil
}
