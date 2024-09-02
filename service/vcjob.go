package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"k8s-plantform/config"

	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	volcanoClient "volcano.sh/apis/pkg/client/clientset/versioned"
	volcanov1alpha1 "volcano.sh/apis/pkg/apis/batch/v1alpha1"
)

var Vcjob vcjob

type vcjob struct{}

// 定义列表的返回内容,Items是pod元素列表,Total是元素数量
type VcjobsResp struct {
	Total int          `json:"total"`
	Items []corev1.Pod `json:"items"`
}
type VcjobNp struct {
	Namespace string `json:"namespace"`
	VcjobNum    int    `json:"vcjob_num"`
}

// 获取vcjob列表,支持过滤,排序,分页
func (p *pod) GetVcjobs(filterName, namespace string, limit, page int) (vcjobsResp *VcjobsResp, err error) {
	//context.TODO()  用于声明一个空的context上下文,用于List方法内设置这个请求超时
	//metav1.ListOptions{} 用于过滤List数据,如label,field等
	vcjobList, err := K8s.volcanoClientSet.BatchV1alpha1().Jobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取vcjob列表失败," + err.Error())
		// 返回给上一层,最终返回给前端,前端捕获到后打印出来
		return nil, errors.New("获取vcjob列表失败," + err.Error())
	}
	// 实例化dataselector结构体,组装数据
	selectableData := &dataSelector{
		GenericDataList: p.toCells(vcjobList.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Name: filterName},
			Paginate: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	// 先过滤
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	// 排序和分页
	data := filtered.Sort().Paginate()
	println("Pod total: ", total)
	// 将DataCell类型转成Pod
	pods := p.fromCells(data.GenericDataList)
	return &VcjobsResp{
		Items: pods,
		Total: total,
	}, nil
}

// 类型转换方法corev1.Pod --> DataCell,DataCell-->corev1.Pod
func (vc *vcjob) toCells(vcjobs []volcanov1alpha1.Job) []DataCell {
	cells := make([]DataCell, len(vcjobs))
	for i := range vcjobs {
		cells[i] = vcjobCell(vcjobs[i])
	}
	return cells
}

func (vc *vcjob) fromCells(cells []DataCell) []volcanov1alpha1.Job {
	vcjobs := make([]volcanov1alpha1.Job, len(cells))
	for i := range cells {
		// cells[i].(podCell)是将DataCell类型转换成podCell
		vcjobs[i] = volcanov1alpha1.Job(cells[i].(vcjobCell))
	}
	return vcjobs
}

// 获取vcjob详情
func (vc *vcjob) GetVcjobDetail(vcjobName, namespace string) (pod *volcanov1alpha1.Job, err error) {
	vcjob, err := K8s.volcanoClientSet.BatchV1alpha1().Jobs(namespace).Get(context.TODO(), vcjobName, metav1.GetOptions{})
	if err != nil {
		logger.Error("获取vcjob详情失败," + err.Error())
		return nil, errors.New("获取vcjob详情失败," + err.Error())
	}

	return vcjob, nil
}

// 删除vcjob
func (vc *vcjob) DeleteVcjob(vcjobName string, namespace string) (err error) {
	err = K8s.volcanoClientSet.BatchV1alpha1().Jobs(namespace).Delete(context.TODO(), vcjobName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error("删除vcjob详情失败," + err.Error())
		return errors.New("删除vcjob详情失败," + err.Error())
	}
	return nil
}

// 更新vcjob
func (vc *vcjob) Updatevcjob(namespace, content string) (err error) {
	var vcjob = &volcanov1alpha1.Job{}
	// 反序列化为Pod对象
	err = json.Unmarshal([]byte(content), vcjob)
	if err != nil {
		logger.Error("反序列化失败," + err.Error())
		return errors.New("反序列化失败," + err.Error())
	}
	// 更新pod
	_, err = K8s.volcanoClientSet.BatchV1alpha1().Jobs(namespace).Update(context.TODO(), vcjob, metav1.UpdateOptions{})
	if err != nil {
		logger.Error("更新vcjob失败," + err.Error())
		return errors.New("更新vcjob失败," + err.Error())
	}
	return nil
}

// 获取vcjob中的任务
func (vc *vcjob) GetVcJobTask(vcjobName string, namespace string) (tasks []volcanov1alpha1.TaskSpec, err error) {
	vcjob, err := vc.GetVcjobDetail(vcjobName, namespace)

	if err != nil {
		return nil, err
	}
	for _, task := range vcjob.Spec.Tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// 获取Pod内容器日志
func (p *pod) GetPodLog(containerName string, podName string, namespace string) (log string, err error) {
	//设置日志配置,容器名,获取内容的配置
	lineLimit := int64(config.PodLogTailLine)
	option := &corev1.PodLogOptions{
		Container: containerName,
		TailLines: &lineLimit,
	}
	// 获取一个request实例
	req := K8s.ClientSet.CoreV1().Pods(namespace).GetLogs(podName, option)
	// 发起stream连接,获取到Response.body
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		logger.Error("更新Pod失败," + err.Error())
		return "", errors.New("更新Pod失败," + err.Error())
	}
	defer podLogs.Close()
	// 将Response.body 写入到缓存区,目的为了转换成string类型
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		logger.Error("复制podLog失败," + err.Error())
		return "", errors.New("复制podLog失败," + err.Error())
	}
	return buf.String(), nil
}

// 获取每个namespace中pod的数量
func (p *pod) GetVcJobNumPerNp() (podsNps []*PodsNp, err error) {
	//获取namespace列表
	namespaceList, err := K8s.ClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, namespace := range namespaceList.Items {
		//获取pod列表
		podList, err := K8s.ClientSet.CoreV1().Pods(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		//组装数据
		podsNp := &PodsNp{
			Namespace: namespace.Name,
			PodNum:    len(podList.Items),
		}
		//添加到podsNps数组中
		podsNps = append(podsNps, podsNp)
	}
	return podsNps, nil
}
