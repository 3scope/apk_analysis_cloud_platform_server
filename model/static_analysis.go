package model

import "gorm.io/gorm"

type StaticAnalysis struct {
	gorm.Model
	AppName     string `json:"appName" gorm:"column:app_name" form:"appName"`
	AppVersion  string `json:"appVersion" gorm:"column:app_version" form:"appVersion"`
	ApkSize     uint   `json:"apkSize" gorm:"colunm:apk_size" form:"apkSize"`
	PackageName string `json:"packageName" gorm:"column:package_name" form:"packageName"`
	MD5         string `json:"md5" gorm:"column:md5" form:"md5"`
	SHA160      string `json:"sha160" gorm:"column:sha160" form:"sha160"`
	SHA256      string `json:"sha256" gorm:"column:sha256"`
	SDK         string `json:"sdk" gorm:"column:sdk" form:"sdk"`
	Certificate string `json:"certificate" gorm:"colunm:certificate" form:"certificate"`
	Email       string `json:"email" gorm:"column:email" form:"email"`
	IPAdress    string `json:"ipAdress" gorm:"column:ip_adress" form:"ipAdress"`
	DomainName  string `json:"domainName" gorm:"colunm:domain_name" form:"domainName"`
	Permission  string `json:"permission" gorm:"column:permission" form:"permission"`

	// Logical foreign key.
	CaseID uint `json:"caseID" gorm:"column:case_id" form:"caseID"`
}
