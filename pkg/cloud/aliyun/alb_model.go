package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alb"
)

type SLBClient struct {
	client *alb.Client
}

// https://help.aliyun.com/document_detail/213629.html 请求参数详细文档

type ServerList struct {
	ServerId    string // 服务器为阿里云的实例取值,如果是ip 为ip
	ServerIp    string // 后端服务的IP
	Port        uint   // 默认是80
	ServerType  string // 后端服务器的类型 Ecs|Eni|Eci|Ip
	Description string // 后端服务器的描述 /^([^\x00-\xff]|[\w.,;/@-]){2,256}$/
	Weight      uint   // 后端服务器的权重 0～ 100 之间,默认是100
	ClientToken string // 保证每次请求的幂等
	DryRun      bool   // 是否预检本次请求，不创建资源，只检测必须参数
}
