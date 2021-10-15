package util

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/sanscope/apk_analysis_cloud_platform_server/modify_log"
)

// To set the time string format.
const TimeFormat = "2006-01-02 15:04:05"

// To set the time location to CST.
var CST = time.FixedZone("CST", 8*60*60)

// The Format function can convert the time.Time object to string.
func GetTimestampNow() string {
	now := time.Now().In(CST).Format(TimeFormat)
	return now
}

func FormatTimestamp(timeString string) string {
	result, err := time.ParseInLocation(TimeFormat, timeString, CST)
	if err != nil {
		log.Panicln(err)
	}
	return result.In(CST).Format(TimeFormat)
}

// To Check the pagination arguments.
func PaginationCheck(pageNumber, pageSize, numberLimit, sizeLimit int) (int, int) {
	switch {
	case pageNumber < 1:
		pageNumber = 1
	case pageNumber > numberLimit:
		pageNumber = numberLimit
	}
	switch {
	case pageSize < 1:
		pageSize = 1
	case pageSize > sizeLimit:
		pageSize = sizeLimit
	}
	limit := pageSize
	offset := (pageNumber - 1) * pageSize
	return limit, offset
}

// Sort in descending order of creation time in default place.
func Sort(sortString string) string {
	if sortString != "" {
		return sortString
	} else {
		return "created_at desc"
	}
}

// The hash method for string.
func GetSHA256FromString(data string) string {
	h := sha256.New()
	io.WriteString(h, data)
	hashString := fmt.Sprintf("%x", h.Sum(nil))
	return hashString
}

// The hash method for file.
func GetSHA256FromFile(path string) string {
	file, err := os.Open(path)
	defer func() {
		if err := file.Close(); err != nil {
			log.Panicln(err)
		}
	}()
	if err != nil {
		log.Panicln(err)
	}
	h := sha256.New()
	_, err = io.Copy(h, file)
	if err != nil {
		log.Panicln(err)
	}
	hashString := fmt.Sprintf("%x", h.Sum(nil))
	return hashString
}
