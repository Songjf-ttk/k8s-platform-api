package service

import (
	"k8s-plantform/config"

	"github.com/wonderivan/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	volcanoClient "volcano.sh/apis/pkg/client/clientset/versioned"
)

// 用于初始化k8s client
var K8s k8s

type k8s struct {
	ClientSet *kubernetes.Clientset
	volcanoClientSet *volcanoClient.Clientset
}

// 初始化k8s
func (k *k8s) Init() {
	// 将kuberconfig文件转换成rest.config对象
	conf, err := clientcmd.BuildConfigFromFlags("", config.Kubeconfig)
	if err != nil {
		panic("获取K8s配置失败:" + err.Error())
	} else {
		logger.Info("获取K8s配置 成功!")
	}
	// 根据rest.config对象,new一个clientset出来
	clientset, err := kubernetes.NewForConfig(conf)
	if err != nil {
		panic("创建K8s client失败:" + err.Error())
	} else {
		logger.Info("创建K8s client 成功!")
	}
	volcanoClientSet, err := volcanoClient.NewForConfig(conf)
	if err != nil {
		panic("创建volcano client失败:"+err.Error())
	} else {
		logger.Info("创建volcano client 成功!")
	}
	k.ClientSet = clientset
	k.volcanoClientSet = volcanoClientSet
}
