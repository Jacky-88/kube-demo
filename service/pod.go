package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var Pod pod

type pod struct {
}

type PodResp struct {
	Items []corev1.Pod `json:"items"`
	Total int          `json:"total"`
}

type PodsNp struct {
	Namespace string `json:"namespace"`
	PodNum    int    `json:"pod_num"`
}

func (p *pod) GetPods(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (podResp *PodResp, err error) {
	podList, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error("获取Pod列表失败." + err.Error())
		return nil, errors.New("获取Pod列表失败." + err.Error())
	}
	selectableData := &dataSelector{
		GenericDataList: p.toCell(podList.Items),
		dataSelectQuery: &DataSelectQuery{
			FilterQuery: &FilterQuery{Name: filterName},
			PaginteQuery: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	// 先过滤
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	// 再排序和分页
	data := filtered.Sort().Paginate()

	pods := p.formCells(data.GenericDataList)

	return &PodResp{
		Items: pods,
		Total: total,
	}, nil
}

// 获取 pod 详情
func (p *pod) GetPodDetail(client *kubernetes.Clientset, podName, namespace string) (*corev1.Pod, error) {
	pod, err := client.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		logger.Error("获取Pod详情失败." + err.Error())
		return nil, errors.New("获取Pod详情失败." + err.Error())
	}
	return pod, nil
}

// toCell方法用于将pod类型数组，转换成DataCell类型数组
func (p *pod) toCell(std []corev1.Pod) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = podCell(std[i])
	}
	return cells

}

// fromCells方法用于将DataCell类型数组，转换成pod类型数组
func (p *pod) formCells(cells []DataCell) []corev1.Pod {
	pods := make([]corev1.Pod, len(cells))
	for i := range cells {
		pods[i] = corev1.Pod(cells[i].(podCell))
	}
	return pods
}
