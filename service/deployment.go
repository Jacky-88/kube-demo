package service

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var Deployment deployment

type deployment struct{}

// 定义列表的返回内容，Items是deployment元素列表，Total为deployment元素数量
type DeploymentResp struct {
	Items []appsv1.Deployment `json:"items"`
	Total int                 `json:"total"`
}

// 获取deployment列表，支持过滤和排序、分页
func (d *deployment) GetDeployments(client *kubernetes.Clientset, filterName,
	namespace string, limit, page int) (deploymentResp *DeploymentResp, err error) {
	//获取deployment类型的deployment列表
	deploymentList, err := client.AppsV1().Deployments(namespace).List(metav1.ListOptions{})
}
