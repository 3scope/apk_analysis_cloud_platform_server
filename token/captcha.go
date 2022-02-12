package token

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/dchest/captcha"
)

// The default attributes of captcha.
const (
	DefaultLength = 6
	DefaultWidth  = 200
	DefaultHeight = 100
)

type CaptchaInstance struct {
	CaptchaLength int `json:"captcha_length"`
	CaptchaWidth  int `json:"captcha_width"`
	CaptchaHeight int `json:"captcha_height"`
}

func (rc *RedisCache) Set(id string, digits []byte) {
	err := rc.SetString(id, string(digits))
	if err != nil {
		fmt.Println(err)
	}
}

func (rc *RedisCache) Get(id string, clear bool) (digits []byte) {
	imageData, err := rc.GetString(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	return []byte(imageData)
}

// To get one captcha image. "imageData" store the base64 data after encoding.
// The return values(captchaId, imageData) are provided to front end.
// The "NewLen" method set the answer of captcha into redis(key: captchaId, value: answer).
func (c *CaptchaInstance) GetCaptcha(redisCache *RedisCache) (string, string, error) {
	// To verify whether the parameters are right.
	captchaData := verifyParameters(c)
	// To set redis to store the captcha answer.
	captcha.SetCustomStore(redisCache)
	// The "NewLen" method is to generate the random numbers and set answer(byte slice) into redis.
	captchaId := captcha.NewLen(captchaData.CaptchaLength)
	// The "imageByte" variable stores the data of captcha image.
	var imageByte bytes.Buffer
	err := captcha.WriteImage(&imageByte, captchaId, captchaData.CaptchaWidth, captchaData.CaptchaHeight)
	// Encode the image into base64 string.
	imageData := base64.StdEncoding.EncodeToString(imageByte.Bytes())
	if err != nil {
		return "", "", err
	}
	return captchaId, imageData, nil
}

// Verify whether the captcha answer is true.
// When is true, delete the key of redis.
func (c *CaptchaInstance) VerifyCaptcha(redisCache *RedisCache, captchaId string, digits string) bool {
	ok := captcha.VerifyString(captchaId, digits)
	if !ok {
		return false
	}
	go func() {
		redisCache.DEL(captchaId)
	}()
	return true
}

func verifyParameters(captcha *CaptchaInstance) *CaptchaInstance {
	reply := captcha
	if captcha.CaptchaLength <= 0 {
		reply.CaptchaLength = DefaultLength
	}
	if captcha.CaptchaWidth <= 0 {
		reply.CaptchaWidth = DefaultWidth
	}
	if captcha.CaptchaHeight <= 0 {
		reply.CaptchaHeight = DefaultHeight
	}
	return reply
}
