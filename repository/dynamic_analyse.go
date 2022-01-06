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
	// The Where object is a struct pointer.
	// Get the entity pointer.
	dynamicAnalysis := GetDynamicAnalysisInstance(request.Entity)
	if dynamicAnalysis == nil {
		return 0, errors.New("invalid type of entity")
	}
	// Query conditions can be added here.
	db = db.Debug().Table("dynamic_analysis").Where(dynamicAnalysis)

	if err := db.Count(&total).Error; err != nil {
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

	// Get the entity pointer.
	dynamicAnalysis := GetDynamicAnalysisInstance(request.Entity)
	if dynamicAnalysis == nil {
		return nil, errors.New("invalid type of entity")
	}
	// The entity is pointer.
	db = db.Debug().Table("dynamic_analysis").Where(dynamicAnalysis)

	if err := db.Order("id ASC").Limit(limit).Offset(offset).Find(&dynamicAnalysiss).Error; err != nil {
		return nil, err
	}
	return dynamicAnalysiss, nil
}

func (repo *DynamicAnalysisRepository) Add(request *Request) (*model.DynamicAnalysis, error) {
	// To judge whether the request entity is existed.
	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("dynamicAnalysis already exists")
	}

	// Get the entity pointer.
	dynamicAnalysis := GetDynamicAnalysisInstance(request.Entity)
	if dynamicAnalysis == nil {
		return nil, errors.New("invalid type of entity")
	}

	// The "dynamic_analysis" variable is pointer.
	if err := repo.DB.Debug().Table("dynamic_analysis").Create(dynamicAnalysis).Error; err != nil {
		return nil, err
	}
	return dynamicAnalysis, nil
}

func (repo *DynamicAnalysisRepository) GetOne(request *Request) (*model.DynamicAnalysis, error) {
	// Get only one instance.
	db := repo.DB
	var count int64

	dynamicAnalysis := GetDynamicAnalysisInstance(request.Entity)
	if dynamicAnalysis == nil {
		return nil, errors.New("invalid type of entity")
	}
	db = db.Debug().Table("dynamic_analysis").Where(dynamicAnalysis)

	if err := db.Limit(1).Find(dynamicAnalysis).Count(&count).Error; err != nil {
		return nil, err
	}
	// To judge whether there are result.
	if count == 0 {
		return nil, nil
	}
	return dynamicAnalysis, nil
}

func (repo *DynamicAnalysisRepository) IsExist(request *Request) (int64, error) {
	var count int64

	// Get the entity pointer.
	dynamicAnalysis := GetDynamicAnalysisInstance(request.Entity)
	if dynamicAnalysis == nil {
		return 0, errors.New("invalid type of entity")
	}

	// The "Entity" is pointer.
	if err := repo.DB.Debug().Table("dynamic_analysis").Where(dynamicAnalysis, "dynamicAnalysisname").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *DynamicAnalysisRepository) Delete(request *Request) (bool, error) {
	dynamicAnalysis := GetDynamicAnalysisInstance(request.Entity)
	if dynamicAnalysis == nil {
		return false, errors.New("invalid type of entity")
	}
	entity, _ := request.Entity.(*DynamicAnalysisEntity)
	dynamicAnalysis.ID = entity.DynamicAnalysisID

	if err := repo.DB.Debug().Table("dynamic_analysis").Where(dynamicAnalysis).Delete(dynamicAnalysis).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *DynamicAnalysisRepository) Update(request *Request) (*model.DynamicAnalysis, error) {
	dynamicAnalysis := GetDynamicAnalysisInstance(request.Entity)
	if dynamicAnalysis == nil {
		return nil, errors.New("invalid type of entity")
	}
	entity, _ := request.Entity.(*DynamicAnalysisEntity)
	dynamicAnalysis.ID = entity.DynamicAnalysisID

	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("dynamicAnalysis already exists")
	}

	// The verification task is handed over to the front end.
	// Cannot change password use this function.
	// Use scan to store the result.
	if err := repo.DB.Debug().Table("dynamic_analysis").Where("id = ?", dynamicAnalysis.ID).Updates(model.DynamicAnalysis{
		AnalysisData: dynamicAnalysis.AnalysisData,
		VideoID:      dynamicAnalysis.VideoID,
		CaseID:       dynamicAnalysis.CaseID,
	}).Scan(&dynamicAnalysis).Error; err != nil {
		return nil, err
	}

	return dynamicAnalysis, nil
}

// To get a entity instance.
func GetDynamicAnalysisInstance(entity interface{}) *model.DynamicAnalysis {
	var dynamicAnalysis model.DynamicAnalysis
	if entity, ok := entity.(*DynamicAnalysisEntity); !ok {
		return nil
	} else {
		dynamicAnalysis = (*entity).DynamicAnalysis
	}
	return &dynamicAnalysis
}
