package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"k8s-demo-test/service"
	"net/http"
)

var Deployment deployment

type deployment struct{}

// 获取deployment列表，支持过滤、排序、分页
func (d *deployment) GetDeployments(ctx *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
		Page       int    `form:"page"`
		Limit      int    `form:"limit"`
		Cluster    string `form:"cluster"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败." + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	fmt.Println(client)

}
