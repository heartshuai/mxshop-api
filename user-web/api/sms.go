package api

import (
	"fmt"
	"math/rand"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// 生成width的验证码
func GenerateSmsCode(witdh int) string {
	numberic := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numberic)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < witdh; i++ {
		fmt.Fprintf(&sb, "%d", numberic[rand.Intn(r)])
	}
	return sb.String()
}

/*
*
发送验证码
*/
func SendSms(ctx *gin.Context) {
	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBind(&sendSmsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", global.ServerConfig.AliSmsInfo.ApiKey, global.ServerConfig.AliSmsInfo.ApiSecret)
	if err != nil {
		panic(err)
	}
	mobile := sendSmsForm.Mobile
	smsCode := GenerateSmsCode(6)
	request := dysmsapi.CreateSendSmsRequest()
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Method = "POST"
	request.RegionId = "cn-hangzhou"
	request.Scheme = "https"
	request.PhoneNumbers = mobile
	request.SignName = "唐山帅启科技"
	request.TemplateCode = "SMS_205325382"
	request.TemplateParam = "{\"code\":" + smsCode + "}"
	response, err := client.SendSms(request)
	fmt.Println(client.DoAction(request, response))
	if err != nil {
		panic(err)
	}
	fmt.Println("response is %#v\n", response)
	//将验证码保存起来
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
		Password: global.ServerConfig.RedisInfo.Password,
	})
	//fmt.Println(rdb)
	rdb.Set(mobile, smsCode, time.Duration(global.ServerConfig.AliSmsInfo.Expire)*time.Second)
	ctx.JSON(200, gin.H{
		"msg": "发送成功",
	})

}
