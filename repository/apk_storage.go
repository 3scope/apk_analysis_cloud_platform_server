package repository

import (
	"errors"
	"log"

	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
	"github.com/sanscope/apk_analysis_cloud_platform_server/util"
	"gorm.io/gorm"
)

type ApkRepository struct {
	DB *gorm.DB
}

// The parameter passed in can be any attribute value of the apk object.
type ApkRepositoryInterface interface {
	GetTotal(request *Query) (total int64)
	List(request *Query) (apks []model.ApkStorage)
	Add(apk model.ApkStorage) (*model.ApkStorage, error)
	Get(apk model.ApkStorage) *model.ApkStorage
	IsExist(apk model.ApkStorage) *model.ApkStorage
	Delete(apk model.ApkStorage) bool
	Update(apk model.ApkStorage) (*model.ApkStorage, error)
}

func (ar *ApkRepository) GetTotal(request *Query) (total int64) {
	db := ar.DB
	var apks []model.ApkStorage
	if request.Where != "" {
		db = db.Where(request.Where)
	}
	if err := db.Find(&apks).Count(&total).Error; err != nil {
		log.Println(err)
		return 0
	}
	return total
}

func (ar *ApkRepository) List(request *Query) (apks []model.ApkStorage) {
	db := ar.DB
	// First to get the total number of data.
	total := ar.GetTotal(request)
	// Get the correct number of data.
	limit, offset := util.PaginationCheck(request.PageNumber, request.PageSize, int(total), 100)
	if request.Where != "" {
		db = db.Where(request.Where)
	}
	if err := db.Order("id asc").Limit(limit).Offset(offset).Find(&apks).Error; err != nil {
		log.Println(err)
		return nil
	}
	return apks
}

func (ar *ApkRepository) Add(apk model.ApkStorage) (*model.ApkStorage, error) {
	if err := ar.DB.Create(&apk).Error; err != nil {
		return nil, errors.New("apk registration failed")
	}
	return &apk, nil
}

func (ar *ApkRepository) Get(apk model.ApkStorage) *model.ApkStorage {
	if err := ar.DB.Where(&apk).Find(&apk).Error; err != nil {
		log.Println(err)
		return nil
	}
	return &apk
}

func (ar *ApkRepository) IsExist(apk model.ApkStorage) *model.ApkStorage {
	if err := ar.DB.Where(&apk).Find(&apk); err != nil {
		log.Println(err)
		return nil
	}
	return &apk
}

func (ar *ApkRepository) Delete(apk model.ApkStorage) bool {
	if err := ar.DB.Delete(&apk); err != nil {
		log.Println(err)
		return false
	}
	return true
}

// This method can not change the password.
func (ar *ApkRepository) Update(apk model.ApkStorage) (*model.ApkStorage, error) {
	if err := ar.DB.Model(&apk).Updates(model.ApkStorage{
		AppName:    apk.AppName,
		UploadedBy: apk.UploadedBy,
	}).Error; err != nil {
		return nil, errors.New("apk update failed")
	}
	return &apk, nil
}
