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
	// The Where object is a struct pointer.
	// Get the entity pointer.
	report := GetReportInstance(request.Entity)
	if report == nil {
		return 0, errors.New("invalid type of entity")
	}
	// Query conditions can be added here.
	db = db.Debug().Table("reports").Where(report)

	if err := db.Count(&total).Error; err != nil {
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

	// Get the entity pointer.
	report := GetReportInstance(request.Entity)
	if report == nil {
		return nil, errors.New("invalid type of entity")
	}
	// The entity is pointer.
	db = db.Debug().Table("reports").Where(report)

	if err := db.Order("id ASC").Limit(limit).Offset(offset).Find(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}

func (repo *ReportRepository) Add(request *Request) (*model.Report, error) {
	// To judge whether the request entity is existed.
	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("report already exists")
	}

	// Get the entity pointer.
	report := GetReportInstance(request.Entity)
	if report == nil {
		return nil, errors.New("invalid type of entity")
	}

	// The "report" variable is pointer.
	if err := repo.DB.Debug().Table("reports").Create(report).Error; err != nil {
		return nil, err
	}
	return report, nil
}

func (repo *ReportRepository) GetOne(request *Request) (*model.Report, error) {
	// Get only one instance.
	db := repo.DB
	var count int64

	report := GetReportInstance(request.Entity)
	if report == nil {
		return nil, errors.New("invalid type of entity")
	}
	db = db.Debug().Table("reports").Where(report)

	if err := db.Limit(1).Find(report).Count(&count).Error; err != nil {
		return nil, err
	}
	// To judge whether there are result.
	if count == 0 {
		return nil, nil
	}
	return report, nil
}

func (repo *ReportRepository) IsExist(request *Request) (int64, error) {
	var count int64

	// Get the entity pointer.
	report := GetReportInstance(request.Entity)
	if report == nil {
		return 0, errors.New("invalid type of entity")
	}

	// The "Entity" is pointer.
	if err := repo.DB.Debug().Table("reports").Where(report, "reportname").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *ReportRepository) Delete(request *Request) (bool, error) {
	report := GetReportInstance(request.Entity)
	if report == nil {
		return false, errors.New("invalid type of entity")
	}
	entity, _ := request.Entity.(*ReportEntity)
	report.ID = entity.ReportID

	if err := repo.DB.Debug().Table("reports").Where(report).Delete(report).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *ReportRepository) Update(request *Request) (*model.Report, error) {
	report := GetReportInstance(request.Entity)
	if report == nil {
		return nil, errors.New("invalid type of entity")
	}
	entity, _ := request.Entity.(*ReportEntity)
	report.ID = entity.ReportID

	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("report already exists")
	}

	// The verification task is handed over to the front end.
	// Cannot change password use this function.
	// Use scan to store the result.
	if err := repo.DB.Debug().Table("reports").Where("id = ?", report.ID).Updates(model.Report{
		ReportName: report.ReportName,
		CaseID:     report.CaseID,
	}).Scan(&report).Error; err != nil {
		return nil, err
	}

	return report, nil
}

// To get a entity instance.
func GetReportInstance(entity interface{}) *model.Report {
	var report model.Report
	if entity, ok := entity.(*ReportEntity); !ok {
		return nil
	} else {
		report = (*entity).Report
	}
	return &report
}
