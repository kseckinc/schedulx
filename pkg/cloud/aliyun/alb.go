package aliyun

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alb"
)

// InitALB 初始化slb
func InitALB(region, accessKey, accessSecret string) (SLBClient, error) {
	client, err := alb.NewClientWithAccessKey(region, accessKey, accessSecret)
	return SLBClient{
		client: client,
	}, err
}

// CreateServers 添加服务器向已经存在的服务器组 alb
func (c *SLBClient) CreateServers(ctx context.Context, groupId string, servers *[]alb.AddServersToServerGroupServers) (
	*alb.AddServersToServerGroupResponse, error) {
	request := alb.CreateAddServersToServerGroupRequest()
	// 初始化参数
	request.ServerGroupId = groupId
	request.Servers = servers
	request.ClientToken = utils.GetUUID()

	backendServers, err := c.client.AddServersToServerGroup(request)
	return backendServers, err
}

func (c *SLBClient) RemoveServer(ctx context.Context, groupId string, servers *[]alb.RemoveServersFromServerGroupServers) (
	*alb.RemoveServersFromServerGroupResponse, error) {
	// 初始化参数
	request := alb.CreateRemoveServersFromServerGroupRequest()
	request.ServerGroupId = groupId
	request.Servers = servers
	request.ClientToken = utils.GetUUID()

	response, err := c.client.RemoveServersFromServerGroup(request)
	return response, err
}
