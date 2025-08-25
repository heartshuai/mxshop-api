package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/proto"
	"net/http"
)

func HandleGrpcErrorToHttp(err error, ctx *gin.Context) {
	//将grpc的code转换成http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "服务器内部错误",
				})
			case codes.InvalidArgument:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})

			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误",
				})
			}

		}
	}
}
func GetUserList(ctx *gin.Context) {
	ip := "127.0.0.1"
	port := 8080
	//拨号连接用户grpc服务
	userConn, err := grpc.Dial(fmt.Sprintf("%s: %d", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接用户服务失败", "msg", err.Error())
		return
	}
	//调用接口
	UserSrcClient := proto.NewUserClient(userConn)
	rsp, err := UserSrcClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 2,
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 获取用户列表失败", "msg", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		data := make(map[string]interface{})
		data["id"] = value.Id
		data["mobile"] = value.Mobile
		data["name"] = value.NickName
		data["birthday"] = value.BirthDay
		data["gender"] = value.Gender
		result = append(result, data)
	}
	ctx.JSON(http.StatusOK, result)

	zap.S().Debug("获取用户列表页")
}
