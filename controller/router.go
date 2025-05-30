package controller

import "github.com/gin-gonic/gin"

var Router router

type router struct{}

func (r *router) InitApiRouter(router *gin.Engine) {
	router.
		// 登录
		POST("/api/login", Login.Auth)
	//集群列表
	//GET("/api/k8s/clusters", Cluster.GetClusters)
}
