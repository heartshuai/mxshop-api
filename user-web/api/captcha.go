package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	c := base64Captcha.NewCaptcha(driver, store)
	ids, b64s, answer, err := c.Generate()
	if err != nil {
		zap.S().Errorf("验证码获取失败", err.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "验证码获取失败",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": ids,
		"picPath":   b64s,
		"answer":    answer,
	})
}
