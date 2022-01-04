package repository

import (
	"github.com/sanscope/apk_analysis_cloud_platform_server/model"
)

// The query arguments in request.
type Request struct {
	PageNumber int `form:"pageNumber"`
	PageSize   int `form:"pageSize"`
	// Use the entity pointer.
	Entity interface{}
}

type UserEntity struct {
	UserID uint `form:"userID"`
	model.User
}

type CaseEntity struct {
	CaseID uint `form:"caseID"`
	model.Case
}

type StaticAnalysisEntity struct {
	StaticAnalysisID uint `form:"staticAnalysisID"`
	model.StaticAnalysis
}
type DynamicAnalysisEntity struct {
	DynamicAnalysisID uint `form:"dynamicAnalysisID"`
	model.DynamicAnalysis
}

type ReportEntity struct {
	ReportID uint `form:"reportID"`
	model.Report
}

type VideoEntity struct {
	VideoID uint `form:"videoID"`
	model.Video
}

/*
* The custom where query.
 */

// The format is ?where=(id=x username=x realName=x role=x email=x phoneNumber=x description=x appName=x reportName=x videoName=x)
// func DecomposeWhere(whereString string) (*Where, bool) {
// 	w := new(Where)
// 	// To determine whether it is contained in parentheses.
// 	if !strings.HasPrefix(whereString, "(") {
// 		return nil, false
// 	}
// 	whereString = strings.TrimPrefix(whereString, "(")
// 	if !strings.HasSuffix(whereString, ")") {
// 		return nil, false
// 	}
// 	whereString = strings.TrimSuffix(whereString, ")")
// 	// Determine whether the format is correct.
// 	defer func() (*Where, bool) {
// 		if err := recover(); err != nil {
// 			log.Println(err)
// 		}
// 		return nil, false
// 	}()
// 	// To split string into slice.
// 	split := strings.Split(whereString, " ")
// 	whereMap := make(map[string]string, len(split))
// 	for _, value := range split {
// 		temp := strings.Split(value, "=")
// 		whereMap[temp[0]] = temp[1]
// 	}
// 	// To get Where struct fields.
// 	wType := reflect.TypeOf(w)
// 	for i := 0; i < wType.NumField(); i++ {
// 		if value, ok := whereMap[wType.Field(i).Name]; ok {
// 			// To change the value of the struct.
// 			reflect.ValueOf(&w).Field(i).SetString(value)
// 		}
// 	}
// 	return w, true
// }
