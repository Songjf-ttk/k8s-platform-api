package controller

import (
	"k8s-plantform/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var Vcjob vcjob

type vcjob struct{}

// Controller中的方法入参是gin.Context 用于从上下文中获取请求参数及定义响应内容
// 流程: 绑定参数 --> 调用service代码 --> 根据调用结果响应具体内容

// 获取vcjob列表,支持分页,过滤,排序
func (vc *vcjob) GetVcjobs(ctx *gin.Context) {
	// 处理入参
	// 匿名结构体用于定义入参,get请求为form格式,其他为json格式
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
		Limit      int    `form:"limit"`
		Page       int    `form:"page"`
	})
	// form 格式使用Bind方法,json格式使用ShouldBindJson方法
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind绑定参数失败," + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "Bind绑定参数失败," + err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Vcjob.GetVcjobs(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取Vcjob列表成功",
		"data": data,
	})
}

// 获取vcjob详情
func (vc *vcjob) GetVcjobDetail(ctx *gin.Context) {
	// 处理入参
	// 匿名结构体用于定义入参,get请求为form格式,其他为json格式
	params := new(struct {
		VcjobName   string `form:"vcjob_name"`
		Namespace string `form:"namespace"`
	})
	// form 格式使用Bind方法,json格式使用ShouldBindJson方法
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind绑定参数失败," + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "Bind绑定参数失败," + err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Vcjob.GetVcjobDetail(params.VcjobName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取Vcjob列表成功",
		"data": data,
	})
}

// 删除Vcjob
func (vc *vcjob) DeleteVcjob(ctx *gin.Context) {
	// 处理入参
	// 匿名结构体用于定义入参,get请求为form格式,其他为json格式
	params := new(struct {
		VcjobName   string `json:"vcjob_name"`
		Namespace string `json:"namespace"`
	})
	// form 格式使用Bind方法,json格式使用ShouldBindJson方法
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind绑定参数失败," + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "Bind绑定参数失败," + err.Error(),
			"data": nil,
		})
		return
	}
	err := service.Vcjob.DeleteVcjob(params.VcjobName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "删除Vcjob成功",
	})
}

// 更新vcjob
func (vc *vcjob) UpdateVcjob(ctx *gin.Context) {
	params := new(struct {
		Namespace string `json:"namespace"`
		Content   string `json:"content"`
	})
	//PUT请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	err := service.Vcjob.UpdateVcjob(params.Namespace, params.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "更新Vcjob成功",
		"data": nil,
	})
}

// 获取vcjob容器
func (vc *vcjob) GetVcjobTaskName(ctx *gin.Context) {
	params := new(struct {
		VcjobName   string `form:"vcjob_name"`
		Namespace string `form:"namespace"`
	})
	//GET请求，绑定参数方法改为ctx.Bind
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Vcjob.GetVcjobTaskName(params.VcjobName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取Vcjob容器成功",
		"data": data,
	})
}

// 获取每个namespace的vcjob数量
func (vc *vcjob) GetVcjobNumPerNp(ctx *gin.Context) {
	data, err := service.Vcjob.GetVcjobNumPerNp()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取每个namespace的vcjob数量成功",
		"data": data,
	})
}



// package controller

// import (
// 	"k8s-plantform/service"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/wonderivan/logger"
// )

// var Vcjob vcjob

// type vc struct{}

// // Controller中的方法入参是gin.Context 用于从上下文中获取请求参数及定义响应内容
// // 流程: 绑定参数 --> 调用service代码 --> 根据调用结果响应具体内容

// // 获取pod列表,支持分页,过滤,排序
// func (vc *vcjob) GetVcjobs(ctx *gin.Context) {
// 	// 处理入参
// 	// 匿名结构体用于定义入参,get请求为form格式,其他为json格式
// 	params := new(struct {
// 		FilterName string `form:"filter_name"`
// 		Namespace  string `form:"namespace"`
// 		Limit      int    `form:"limit"`
// 		Page       int    `form:"page"`
// 	})
// 	// form 格式使用Bind方法,json格式使用ShouldBindJson方法
// 	if err := ctx.Bind(params); err != nil {
// 		logger.Error("Bind绑定参数失败," + err.Error())
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg":  "Bind绑定参数失败," + err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}
// 	data, err := service.vcjob.GetVcjobs(params.FilterName, params.Namespace, params.Limit, params.Page)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg":  err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg":  "获取Pod列表成功",
// 		"data": data,
// 	})
// }

// // 获取pod详情
// func (vc *vcjob) GetVcjobDetail(ctx *gin.Context) {
// 	// 处理入参
// 	// 匿名结构体用于定义入参,get请求为form格式,其他为json格式
// 	params := new(struct {
// 		PodName   string `form:"pod_name"`
// 		Namespace string `form:"namespace"`
// 	})
// 	// form 格式使用Bind方法,json格式使用ShouldBindJson方法
// 	if err := ctx.Bind(params); err != nil {
// 		logger.Error("Bind绑定参数失败," + err.Error())
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg":  "Bind绑定参数失败," + err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}
// 	data, err := service.vcjob.GetVcjobDetail(params.PodName, params.Namespace)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg":  err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg":  "获取Pod列表成功",
// 		"data": data,
// 	})
// }

// // 删除Pod
// func (vc *vcjob) DeleteVcjob(ctx *gin.Context) {
// 	// 处理入参
// 	// 匿名结构体用于定义入参,get请求为form格式,其他为json格式
// 	params := new(struct {
// 		PodName   string `json:"pod_name"`
// 		Namespace string `json:"namespace"`
// 	})
// 	// form 格式使用Bind方法,json格式使用ShouldBindJson方法
// 	if err := ctx.ShouldBindJSON(params); err != nil {
// 		logger.Error("Bind绑定参数失败," + err.Error())
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg":  "Bind绑定参数失败," + err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}
// 	err := service.vcjob.DeleteVcjob(params.PodName, params.Namespace)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg":  err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "删除Pod成功",
// 	})
// }

// // 更新pod
// func (vc *vcjob) Updatevcjob(ctx *gin.Context) {
// 	params := new(struct {
// 		Namespace string `json:"namespace"`
// 		Content   string `json:"content"`
// 	})
// 	//PUT请求，绑定参数方法改为ctx.ShouldBindJSON
// 	if err := ctx.ShouldBindJSON(params); err != nil {
// 		logger.Error("Bind请求参数失败, " + err.Error())
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg":  err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}
// 	err := service.Pod.Updatevcjob(params.Namespace, params.Content)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg":  err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg":  "更新Pod成功",
// 		"data": nil,
// 	})
// }

// // 获取pod容器
// func (vc *vcjob) GetPodContainer(ctx *gin.Context) {
// 	params := new(struct {
// 		PodName   string `form:"pod_name"`
// 		Namespace string `form:"namespace"`
// 	})
// 	//GET请求，绑定参数方法改为ctx.Bind
// 	if err := ctx.Bind(params); err != nil {
// 		logger.Error("Bind请求参数失败, " + err.Error())
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg":  err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}
// 	data, err := service.Pod.GetPodContainer(params.PodName, params.Namespace)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg":  err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg":  "获取Pod容器成功",
// 		"data": data,
// 	})
// }

// // 获取pod中容器日志
// func (vc *vcjob) GetVcjobLog(ctx *gin.Context) {
// 	params := new(struct {
// 		ContainerName string `form:"container_name"`
// 		PodName       string `form:"pod_name"`
// 		Namespace     string `form:"namespace"`
// 	})
// 	//GET请求，绑定参数方法改为ctx.Bind
// 	if err := ctx.Bind(params); err != nil {
// 		logger.Error("Bind请求参数失败, " + err.Error())
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg":  err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}
// 	data, err := service.Pod.GetVcjobLog(params.ContainerName, params.PodName, params.Namespace)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg":  err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg":  "获取Pod中容器日志成功",
// 		"data": data,
// 	})
// }

// // 获取每个namespace的pod数量
// func (vc *vcjob) GetVcJobNumPerNp(ctx *gin.Context) {
// 	data, err := service.Pod.GetVcJobNumPerNp()
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg":  err.Error(),
// 			"data": nil,
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg":  "获取每个namespace的pod数量成功",
// 		"data": data,
// 	})
// }
