package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/certificate"
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

	pods :=
}

// toCell方法用于将pod类型数组，转换成DataCell类型数组
func (p *pod) toCell(std []corev1.Pod) []DataCell {
	cells := make([]DataCell,len(std))
	for i := range std {
		cells[i] = podCell(std[i])
	}
	return cells

}
