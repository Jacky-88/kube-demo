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

func (p *pod) GetPods(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (podResp *PodResp, err) {
	podList, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error("获取Pod列表失败." + err.Error())
		return nil, errors.New("获取Pod列表失败." + err.Error())
	}
	selectableData := &dataSelector{
		GenericDataList: p.to
	}
}
