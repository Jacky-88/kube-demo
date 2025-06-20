package controller

import "github.com/gin-gonic/gin"

var Router router

type router struct{}

func (r *router) InitApiRouter(router *gin.Engine) {
	router.
		// 登录
		POST("/api/login", Login.Auth).
		//pod 操作
		GET("/api/k8s/pods", Pod.GetPods).
		GET("/api/k8s/pod/detail", Pod.GetPodDetail).

		// deployment 操作
		GET("/api/k8s/deployments")
}
