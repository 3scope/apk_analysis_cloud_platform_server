package repository

import (
	"errors"
	"log"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/util"
	"gorm.io/gorm"
)

type ReportRepository struct {
	DB *gorm.DB
}

// The parameter passed in can be any attribute value of the report object.
type ReportRepositoryInterface interface {
	GetTotal(request *Query) (total int64)
	List(request *Query) (reports []model.ReportStorage)
	Add(report model.ReportStorage) (*model.ReportStorage, error)
	Get(report model.ReportStorage) *model.ReportStorage
	IsExist(report model.ReportStorage) *model.ReportStorage
	Delete(report model.ReportStorage) bool
	Update(report model.ReportStorage) (*model.ReportStorage, error)
}

func (rr *ReportRepository) GetTotal(request *Query) (total int64) {
	db := rr.DB
	var reports []model.ReportStorage
	if request.Where != "" {
		db = db.Where(request.Where)
	}
	if err := db.Find(&reports).Count(&total).Error; err != nil {
		log.Println(err)
		return 0
	}
	return total
}

func (rr *ReportRepository) List(request *Query) (reports []model.ReportStorage) {
	db := rr.DB
	// First to get the total number of data.
	total := rr.GetTotal(request)
	// Get the correct number of data.
	limit, offset := util.PaginationCheck(request.PageNumber, request.PageSize, int(total), 100)
	if request.Where != "" {
		db = db.Where(request.Where)
	}
	if err := db.Order("id asc").Limit(limit).Offset(offset).Find(&reports).Error; err != nil {
		log.Println(err)
		return nil
	}
	return reports
}

func (rr *ReportRepository) Add(report model.ReportStorage) (*model.ReportStorage, error) {
	if err := rr.DB.Create(&report).Error; err != nil {
		return nil, errors.New("report registration failed")
	}
	return &report, nil
}

func (rr *ReportRepository) Get(report model.ReportStorage) *model.ReportStorage {
	if err := rr.DB.Where(&report).Find(&report).Error; err != nil {
		log.Println(err)
		return nil
	}
	return &report
}

func (rr *ReportRepository) IsExist(report model.ReportStorage) *model.ReportStorage {
	if err := rr.DB.Where(&report).Find(&report); err != nil {
		log.Println(err)
		return nil
	}
	return &report
}

func (rr *ReportRepository) Delete(report model.ReportStorage) bool {
	if err := rr.DB.Delete(&report); err != nil {
		log.Println(err)
		return false
	}
	return true
}

// This method can not change the password.
func (rr *ReportRepository) Update(report model.ReportStorage) (*model.ReportStorage, error) {
	if err := rr.DB.Model(&report).Updates(model.ReportStorage{
		ReportName: report.ReportName,
		UploadedBy: report.UploadedBy,
		AppName:    report.AppName,
	}).Error; err != nil {
		return nil, errors.New("report update failed")
	}
	return &report, nil
}
