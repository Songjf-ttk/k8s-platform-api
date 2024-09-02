package config

import "time"

const (
	//监听地址
	ListenAddr = "0.0.0.0:9091"
	//kubeconfig路径
	Kubeconfig = "config/cka"
	//pod日志tail显示行数
	PodLogTailLine = 2000
	//登录账号密码
	AdminUser = "admin"
	AdminPwd  = "123456"

	//数据库配置
	DbType = "mysql"
	DbHost = "127.0.0.1"
	DbPort = 3306
	DbName = "k8s_dashboard"
	DbUser = "root"
	DbPwd  = "123456"
	//打印mysql debug sql日志
	LogMode = false
	//连接池配置
	MaxIdleConns = 10               //最大空闲连接
	MaxOpenConns = 100              //最大连接数
	MaxLifeTime  = 30 * time.Second //最大生存时间
)
