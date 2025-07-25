package service

import (
	appsv1 "k8s.io/api/apps/v1"
)

var Deployment deployment

type deployment struct{}

// 定义列表的返回内容，Items是deployment元素列表，Total为deployment元素数量
type DeploymentResp struct {
	Items []appsv1.Deployment `json:"items"`
	Total int                 `json:"total"`
}

// 获取deployment列表，支持过滤和排序、分页
//func (d *deployment) GetDeployments(client *kubernetes.Clientset, filterName,
//	namespace string, limit, page int) (deploymentResp *DeploymentResp, err error) {
//	//获取deployment类型的deployment列表
//	deploymentList, err := client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
//	if err != nil {
//		logger.Error("获取Deployment列表失败." + err.Error())
//		return nil, errors.New("获取Deployment列表失败." + err.Error())
//	}
//	//将deploymentList中的deployment列表(Items)，放进dataselector对象中，进行排序
//	selectableData := &dataSelector{
//		GenericDataList: d.toCells(deploymentList.Items),
//	}
//}
