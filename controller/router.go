package controller

import (
	"github.com/gin-gonic/gin"
)

// 初始化router类型的对象,首字母大写,用于跨包调用
var Router router

// 声明一个router的结构体
type router struct{}

func (r *router) InitApiRouter(router *gin.Engine) {
	router.
		// login
		POST("/api/login", Login.Auth).
		// vcjobs
		GET("/api/k8s/vcjobs", Vcjob.GetVcjobs).
		GET("/api/k8s/vcjob/detail", Vcjob.GetVcjobDetail).
		POST("/api/k8s/vcjobs", Vcjob.DeleteVcjob).
		DELETE("/api/k8s/vcjob/del", Vcjob.DeleteVcjob).
		PUT("/api/k8s/vcjob/create", Vcjob.CreateVcjob).
		PUT("/api/k8s/vcjob/update", Vcjob.UpdateVcjob).
		GET("/api/k8s/vcjob/taskname", Vcjob.GetVcjobTaskName).
		GET("/api/k8s/vcjob/numnp", Vcjob.GetVcjobNumPerNp).
		
		// Pods
		GET("/api/k8s/pods", Pod.GetPods).
		GET("/api/k8s/pod/detail", Pod.GetPodDetail).
		POST("/api/k8s/pods", Pod.DeletePod).
		DELETE("/api/k8s/pod/del", Pod.DeletePod).
		PUT("/api/k8s/pod/update", Pod.UpdatePod).
		GET("/api/k8s/pod/container", Pod.GetPodContainer).
		GET("/api/k8s/pod/log", Pod.GetPodLog).
		GET("/api/k8s/pod/numnp", Pod.GetPodNumPerNp).
		//deployment操作
		GET("/api/k8s/deployments", Deployment.GetDeployments).
		GET("/api/k8s/deployment/detail", Deployment.GetDeploymentDetail).
		PUT("/api/k8s/deployment/scale", Deployment.ScaleDeployment).
		DELETE("/api/k8s/deployment/del", Deployment.DeleteDeployment).
		PUT("/api/k8s/deployment/restart", Deployment.RestartDeployment).
		PUT("/api/k8s/deployment/update", Deployment.UpdateDeployment).
		GET("/api/k8s/deployment/numnp", Deployment.GetDeployNumPerNp).
		POST("/api/k8s/deployment/create", Deployment.CreateDeployment).
		//daemonset操作
		GET("/api/k8s/daemonsets", DaemonSet.GetDaemonSets).
		GET("/api/k8s/daemonset/detail", DaemonSet.GetDaemonSetDetail).
		DELETE("/api/k8s/daemonset/del", DaemonSet.DeleteDaemonSet).
		PUT("/api/k8s/daemonset/update", DaemonSet.UpdateDaemonSet).
		//statefulset操作
		GET("/api/k8s/statefulsets", StatefulSet.GetStatefulSets).
		GET("/api/k8s/statefulset/detail", StatefulSet.GetStatefulSetDetail).
		DELETE("/api/k8s/statefulset/del", StatefulSet.DeleteStatefulSet).
		PUT("/api/k8s/statefulset/update", StatefulSet.UpdateStatefulSet).
		//node操作
		GET("/api/k8s/nodes", Node.GetNodes).
		GET("/api/k8s/node/detail", Node.GetNodeDetail).
		//namespace操作
		GET("/api/k8s/namespaces", Namespace.GetNamespaces).
		GET("/api/k8s/namespace/detail", Namespace.GetNamespaceDetail).
		DELETE("/api/k8s/namespace/del", Namespace.DeleteNamespace).
		POST("/api/k8s/namespace/create", Namespace.CreateNamespace).
		//pv操作
		GET("/api/k8s/pvs", Pv.GetPvs).
		GET("/api/k8s/pv/detail", Pv.GetPvDetail).
		//service操作
		GET("/api/k8s/services", Servicev1.GetServices).
		GET("/api/k8s/service/detail", Servicev1.GetServiceDetail).
		DELETE("/api/k8s/service/del", Servicev1.DeleteService).
		PUT("/api/k8s/service/update", Servicev1.UpdateService).
		POST("/api/k8s/service/create", Servicev1.CreateService).
		//ingress操作
		GET("/api/k8s/ingresses", Ingress.GetIngresses).
		GET("/api/k8s/ingress/detail", Ingress.GetIngressDetail).
		DELETE("/api/k8s/ingress/del", Ingress.DeleteIngress).
		PUT("/api/k8s/ingress/update", Ingress.UpdateIngress).
		POST("/api/k8s/ingress/create", Ingress.CreateIngress).
		//pvc操作
		GET("/api/k8s/pvcs", Pvc.GetPvcs).
		GET("/api/k8s/pvc/detail", Pvc.GetPvcDetail).
		DELETE("/api/k8s/pvc/del", Pvc.DeletePvc).
		PUT("/api/k8s/pvc/update", Pvc.UpdatePvc).
		// secret
		GET("/api/k8s/secrets", Secret.GetSecrets).
		GET("/api/k8s/secret/detail", Secret.GetSecretDetail).
		DELETE("/api/k8s/secret/del", Secret.DeleteSecret).
		PUT("/api/k8s/secret/update", Secret.UpdateSecret).
		// workflows
		// GET("/api/k8s/workflows", Workflow.GetList).
		// GET("/api/k8s/workflow/detail", Workflow.GetById).
		// POST("/api/k8s/workflow/create", Workflow.Create).
		// DELETE("/api/k8s/workflow/del", Workflow.DelById).
		// Configmaps
		GET("/api/k8s/configmaps", ConfigMap.GetConfigMaps).
		GET("/api/k8s/configmap/detail", ConfigMap.GetConfigMapDetail).
		DELETE("/api/k8s/configmap/del", ConfigMap.DeleteConfigMap).
		PUT("/api/k8s/configmap/update", ConfigMap.UpdateConfigMap)
}
