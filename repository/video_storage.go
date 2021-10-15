package repository

import (
	"errors"
	"log"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/util"
	"gorm.io/gorm"
)

type VideoRepository struct {
	DB *gorm.DB
}

// The parameter passed in can be any attribute value of the video object.
type VideoRepositoryInterface interface {
	GetTotal(request *Query) (total int64)
	List(request *Query) (videos []model.VideoStorage)
	Add(video model.VideoStorage) (*model.VideoStorage, error)
	Get(video model.VideoStorage) *model.VideoStorage
	IsExist(video model.VideoStorage) *model.VideoStorage
	Delete(video model.VideoStorage) bool
	Update(video model.VideoStorage) (*model.VideoStorage, error)
}

func (vr *VideoRepository) GetTotal(request *Query) (total int64) {
	db := vr.DB
	var videos []model.VideoStorage
	if request.Where != "" {
		db = db.Where(request.Where)
	}
	if err := db.Find(&videos).Count(&total).Error; err != nil {
		log.Println(err)
		return 0
	}
	return total
}

func (vr *VideoRepository) List(request *Query) (videos []model.VideoStorage) {
	db := vr.DB
	// First to get the total number of data.
	total := vr.GetTotal(request)
	// Get the correct number of data.
	limit, offset := util.PaginationCheck(request.PageNumber, request.PageSize, int(total), 100)
	if request.Where != "" {
		db = db.Where(request.Where)
	}
	if err := db.Order("id asc").Limit(limit).Offset(offset).Find(&videos).Error; err != nil {
		log.Println(err)
		return nil
	}
	return videos
}

func (vr *VideoRepository) Add(video model.VideoStorage) (*model.VideoStorage, error) {
	if err := vr.DB.Create(&video).Error; err != nil {
		return nil, errors.New("video registration failed")
	}
	return &video, nil
}

func (vr *VideoRepository) Get(video model.VideoStorage) *model.VideoStorage {
	if err := vr.DB.Where(&video).Find(&video).Error; err != nil {
		log.Println(err)
		return nil
	}
	return &video
}

func (vr *VideoRepository) IsExist(video model.VideoStorage) *model.VideoStorage {
	if err := vr.DB.Where(&video).Find(&video); err != nil {
		log.Println(err)
		return nil
	}
	return &video
}

func (vr *VideoRepository) Delete(video model.VideoStorage) bool {
	if err := vr.DB.Delete(&video); err != nil {
		log.Println(err)
		return false
	}
	return true
}

// This method can not change the password.
func (vr *VideoRepository) Update(video model.VideoStorage) (*model.VideoStorage, error) {
	if err := vr.DB.Model(&video).Updates(model.VideoStorage{
		VideoName:  video.VideoName,
		UploadedBy: video.UploadedBy,
		AppName:    video.AppName,
	}).Error; err != nil {
		return nil, errors.New("video update failed")
	}
	return &video, nil
}
