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
	// The Where object is a struct pointer.
	// Get the entity pointer.
	caseInstance := GetCaseInstance(request.Entity)
	if caseInstance == nil {
		return 0, errors.New("invalid type of entity")
	}
	// Query conditions can be added here.
	db = db.Debug().Table("cases").Where(caseInstance)

	if err := db.Count(&total).Error; err != nil {
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

	// Get the entity pointer.
	caseInstance := GetCaseInstance(request.Entity)
	if caseInstance == nil {
		return nil, errors.New("invalid type of entity")
	}
	// The entity is pointer.
	db = db.Debug().Table("cases").Where(caseInstance)

	if err := db.Order("id ASC").Limit(limit).Offset(offset).Find(&caseInstances).Error; err != nil {
		return nil, err
	}
	return caseInstances, nil
}

func (repo *CaseRepository) Add(request *Request) (*model.Case, error) {
	// To judge whether the request entity is existed.
	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("case instance already exists")
	}

	// Get the entity pointer.
	caseInstance := GetCaseInstance(request.Entity)
	if caseInstance == nil {
		return nil, errors.New("invalid type of entity")
	}

	// The "caseInstance" variable is pointer.
	if err := repo.DB.Debug().Table("cases").Create(caseInstance).Error; err != nil {
		return nil, err
	}
	return caseInstance, nil
}

func (repo *CaseRepository) GetOne(request *Request) (*model.Case, error) {
	// Get only one instance.
	db := repo.DB
	var count int64

	caseInstance := GetCaseInstance(request.Entity)
	if caseInstance == nil {
		return nil, errors.New("invalid type of entity")
	}
	db = db.Debug().Table("cases").Where(caseInstance)

	if err := db.Limit(1).Find(caseInstance).Count(&count).Error; err != nil {
		return nil, err
	}
	// To judge whether there are result.
	if count == 0 {
		return nil, nil
	}
	return caseInstance, nil
}

func (repo *CaseRepository) IsExist(request *Request) (int64, error) {
	var count int64

	// Get the entity pointer.
	caseInstance := GetCaseInstance(request.Entity)
	if caseInstance == nil {
		return 0, errors.New("invalid type of entity")
	}

	// The "Entity" is pointer.
	if err := repo.DB.Debug().Table("cases").Where(caseInstance, "case_name").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *CaseRepository) Delete(request *Request) (bool, error) {
	caseInstance := GetCaseInstance(request.Entity)
	if caseInstance == nil {
		return false, errors.New("invalid type of entity")
	}
	entity, _ := request.Entity.(*CaseEntity)
	caseInstance.ID = entity.CaseID

	if err := repo.DB.Debug().Table("cases").Where(caseInstance).Delete(caseInstance).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *CaseRepository) Update(request *Request) (*model.Case, error) {
	caseInstance := GetCaseInstance(request.Entity)
	if caseInstance == nil {
		return nil, errors.New("invalid type of entity")
	}
	entity, _ := request.Entity.(*CaseEntity)
	caseInstance.ID = entity.CaseID

	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("case instance already exists")
	}

	// The verification task is handed over to the front end.
	// Cannot change password use this function.
	// Use scan to store the result.
	if err := repo.DB.Model(&model.Case{}).Where("id = ?", caseInstance.ID).Updates(model.Case{
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

	return caseInstance, nil
}

// To get a entity instance.
func GetCaseInstance(entity interface{}) *model.Case {
	var caseInstance model.Case
	if entity, ok := entity.(*CaseEntity); !ok {
		return nil
	} else {
		caseInstance = (*entity).Case
	}
	return &caseInstance
}
