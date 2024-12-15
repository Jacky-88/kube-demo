package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var K8s k8s

type k8s struct {
	ClientMap map[string]*kubernetes.Clientset
	KubeConfMap map[string]string
}

func(k *k8s) GetClient(cluster string) (*kubernetes.Clientset,error) {
	client,ok := k.ClientMap[cluster]
	if !ok {
		return nil ,errors.New(fmt.Sprintf("集群：%s 不存在，无法获取client",cluster))
	}
	return client,nil
}

func (k *k8s) Init() {
	mp := map[string]string{}
	k.ClientMap = map[string]*kubernetes.Clientset{}
	if err:= json.Unmarshal([]byte())
}

