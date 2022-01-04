package repository

import (
	"errors"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/util"
	"gorm.io/gorm"
)

type VideoRepository struct {
	DB *gorm.DB
}

// The parameter passed in can be any attribute value of the entity.
type VideoRepositoryInterface interface {
	GetTotal(request *Request) (int64, error)
	List(request *Request) ([]model.Video, error)
	Add(request *Request) (*model.Video, error)
	GetOne(request *Request) (*model.Video, error)
	IsExist(request *Request) (int64, error)
	Delete(request *Request) (bool, error)
	Update(request *Request) (*model.Video, error)
}

func (repo *VideoRepository) GetTotal(request *Request) (int64, error) {
	// The default database.
	var total int64
	db := repo.DB
	var videos []model.Video
	// The Where object is a struct pointer.
	if request.Entity != nil {
		db = db.Where(request.Entity)
	}
	if err := db.Find(&videos).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (repo *VideoRepository) List(request *Request) ([]model.Video, error) {
	videos := make([]model.Video, 0)
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

	if err := db.Order("id asc").Limit(limit).Offset(offset).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func (repo *VideoRepository) Add(request *Request) (*model.Video, error) {
	var video model.Video
	// To judge whether the request entity is existed.
	if count, err := repo.IsExist(request); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.New("video already exists")
	}

	// Use scan to store the result.
	if err := repo.DB.Create(request.Entity).Scan(&video).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

func (repo *VideoRepository) GetOne(request *Request) (*model.Video, error) {
	// Get only one instance.
	var video model.Video
	db := repo.DB

	if request.Entity != nil {
		db = db.Where(request.Entity)
	}

	if err := db.Limit(1).Find(&video).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

func (repo *VideoRepository) IsExist(request *Request) (int64, error) {
	var count int64
	// The "Entity" is pointer.
	if err := repo.DB.Where(request.Entity).Find(&model.Video{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (repo *VideoRepository) Delete(request *Request) (bool, error) {
	if err := repo.DB.Model(&model.Video{}).Where(request.Entity).Delete(request.Entity).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *VideoRepository) Update(request *Request) (*model.Video, error) {
	var video model.Video
	// The verification task is handed over to the front end.
	// Use scan to store the result.
	if err := repo.DB.Model(&model.Video{}).Where(request.Entity).Updates(model.Video{
		VideoName: video.VideoName,
		VideoTime: video.VideoTime,
	}).Scan(&video).Error; err != nil {
		return nil, err
	}

	return &video, nil
}
