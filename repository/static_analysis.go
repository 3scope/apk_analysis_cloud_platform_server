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
	// The Where object is a struct pointer.
	// Get the entity pointer.
	staticAnalysis := GetStaticAnalysisInstance(request.Entity)
	if staticAnalysis == nil {
		return 0, errors.New("invalid type of entity")
	}
	// Query conditions can be added here.
	db = db.Debug().Table("static_analysis").Where(staticAnalysis)

	if err := db.Count(&total).Error; err != nil {
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

	// Get the entity pointer.
	staticAnalysis := GetStaticAnalysisInstance(request.Entity)
	if staticAnalysis == nil {
		return nil, errors.New("invalid type of entity")
	}
	// The entity is pointer.
	db = db.Debug().Table("static_analysis").Where(staticAnalysis)

	if err := db.Order("id ASC").Limit(limit).Offset(offset).Find(&staticAnalysiss).Error; err != nil {
		return nil, err
	}
	return staticAnalysiss, nil
}

func (repo *StaticAnalysisRepository) Add(request *Request) (*model.StaticAnalysis, error) {
	// To judge whether the request entity is existed.
	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("staticAnalysis already exists")
	}

	// Get the entity pointer.
	staticAnalysis := GetStaticAnalysisInstance(request.Entity)
	if staticAnalysis == nil {
		return nil, errors.New("invalid type of entity")
	}

	// The "static_analysis" variable is pointer.
	if err := repo.DB.Debug().Table("static_analysis").Create(staticAnalysis).Error; err != nil {
		return nil, err
	}
	return staticAnalysis, nil
}

func (repo *StaticAnalysisRepository) GetOne(request *Request) (*model.StaticAnalysis, error) {
	// Get only one instance.
	db := repo.DB
	var count int64

	staticAnalysis := GetStaticAnalysisInstance(request.Entity)
	if staticAnalysis == nil {
		return nil, errors.New("invalid type of entity")
	}
	db = db.Debug().Table("static_analysis").Where(staticAnalysis)

	if err := db.Limit(1).Find(staticAnalysis).Count(&count).Error; err != nil {
		return nil, err
	}
	// To judge whether there are result.
	if count == 0 {
		return nil, nil
	}
	return staticAnalysis, nil
}

func (repo *StaticAnalysisRepository) IsExist(request *Request) (int64, error) {
	var count int64

	// Get the entity pointer.
	staticAnalysis := GetStaticAnalysisInstance(request.Entity)
	if staticAnalysis == nil {
		return 0, errors.New("invalid type of entity")
	}

	// The "Entity" is pointer.
	if err := repo.DB.Debug().Table("static_analysis").Where(staticAnalysis, "app_name").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *StaticAnalysisRepository) Delete(request *Request) (bool, error) {
	staticAnalysis := GetStaticAnalysisInstance(request.Entity)
	if staticAnalysis == nil {
		return false, errors.New("invalid type of entity")
	}
	entity, _ := request.Entity.(*StaticAnalysisEntity)
	staticAnalysis.ID = entity.StaticAnalysisID

	if err := repo.DB.Debug().Table("static_analysis").Where(staticAnalysis).Delete(staticAnalysis).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *StaticAnalysisRepository) Update(request *Request) (*model.StaticAnalysis, error) {
	staticAnalysis := GetStaticAnalysisInstance(request.Entity)
	if staticAnalysis == nil {
		return nil, errors.New("invalid type of entity")
	}
	entity, _ := request.Entity.(*StaticAnalysisEntity)
	staticAnalysis.ID = entity.StaticAnalysisID

	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("staticAnalysis already exists")
	}

	// The verification task is handed over to the front end.
	// Cannot change password use this function.
	// Use scan to store the result.
	if err := repo.DB.Debug().Table("static_analysis").Where("id = ?", staticAnalysis.ID).Updates(model.StaticAnalysis{
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

	return staticAnalysis, nil
}

// To get a entity instance.
func GetStaticAnalysisInstance(entity interface{}) *model.StaticAnalysis {
	var staticAnalysis model.StaticAnalysis
	if entity, ok := entity.(*StaticAnalysisEntity); !ok {
		return nil
	} else {
		staticAnalysis = (*entity).StaticAnalysis
	}
	return &staticAnalysis
}
