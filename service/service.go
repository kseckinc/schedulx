package service

import (
	"context"
	"sync"

	"github.com/galaxy-future/schedulx/api/types"
	"github.com/galaxy-future/schedulx/register/config/client"
	"github.com/galaxy-future/schedulx/repository/model/db"

	"github.com/galaxy-future/schedulx/register/config/log"
	"github.com/galaxy-future/schedulx/repository"
)

type ServiceSvc struct {
}

var (
	serviceSvc     *ServiceSvc
	serviceSvcOnce sync.Once
)

type ServiceCreateSvcRequest struct {
	ServiceInfo *types.ServiceInfo `json:"service_info"`
}

type ServiceCreateSvcResponse struct {
	ServiceClusterId int64 `json:"service_cluster_id"`
}

func GetServiceIns() *ServiceSvc {
	serviceSvcOnce.Do(
		func() {
			serviceSvc = &ServiceSvc{}
		})
	return serviceSvc
}

func (s *ServiceSvc) GetServiceList(ctx context.Context, page, pageSize int, serviceName, lang string) (map[string]interface{}, error) {
	var err error
	list, total, err := repository.GetServiceRepoInst().GetServiceList(ctx, page, pageSize, serviceName, lang)
	if err != nil {
		log.Logger.Errorf("repository.GetServiceRepoInst().GetServiceList error:", err)
	}
	ret := map[string]interface{}{
		"service_list": list,
		"pager": struct {
			PageNumber int   `json:"page_number"`
			PageSize   int   `json:"page_size"`
			Total      int64 `json:"total"`
		}{
			PageSize:   pageSize,
			PageNumber: page,
			Total:      total,
		},
	}
	return ret, nil
}

func (s *ServiceSvc) GetExpandHistory(ctx context.Context, page, pageSize, serviceClusterId int) (map[string]interface{}, error) {
	var err error
	tempList, total, err := repository.GetScheduleTemplateRepoInst().GetScheduleTempList(ctx, page, pageSize, serviceClusterId)
	if err != nil {
		log.Logger.Errorf("serviceClusterId:%v page:%v pageSize:%v error:%v", serviceClusterId, page, pageSize, err)
		return nil, err
	}
	ret := map[string]interface{}{
		"schedule_task_list": tempList,
		"pager": struct {
			PageNumber int   `json:"page_number"`
			PageSize   int   `json:"page_size"`
			Total      int64 `json:"total"`
		}{
			PageSize:   pageSize,
			PageNumber: page,
			Total:      total,
		},
	}
	return ret, nil
}

func (s *ServiceSvc) Detail(ctx context.Context, serviceName string) (map[string]interface{}, error) {
	var err error
	serviceDetail, err := repository.GetServiceRepoInst().GetServiceDetail(ctx, serviceName)
	if err != nil {
		return nil, err
	}
	ret := map[string]interface{}{
		"service_info": serviceDetail,
	}
	return ret, nil
}

func (s *ServiceSvc) UpdateDesc(ctx context.Context, serviceName, description string) (map[string]interface{}, error) {
	var err error
	ret, err := repository.GetServiceRepoInst().UpdateDesc(ctx, serviceName, description)
	if err != nil {
		log.Logger.Errorf("update service_name:%v error:%v", serviceName, err)
		return nil, err
	}
	return map[string]interface{}{
		"records": ret,
	}, nil
}

func (s *ServiceSvc) CreateService(ctx context.Context, svcReq *ServiceCreateSvcRequest) (*ServiceCreateSvcResponse, error) {
	var err error
	svcResp := &ServiceCreateSvcResponse{}
	svcObj := &db.Service{
		ServiceName: svcReq.ServiceInfo.ServiceName,
		Description: svcReq.ServiceInfo.Description,
		Language:    svcReq.ServiceInfo.Language,
	}
	dbo := client.WriteDBCli.Begin().WithContext(ctx)
	defer func() {
		if err != nil {
			dbo.Rollback()
			return
		}
		dbo.Commit()
	}()
	err = db.Create(svcObj, dbo)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	// 创建一个服务默认集群
	svcClusterObj := &db.ServiceCluster{
		ServiceName: svcReq.ServiceInfo.ServiceName,
		ClusterName: "default",
		//CreateAt: time.Now(),
	}
	err = db.Create(svcClusterObj, dbo)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	svcResp.ServiceClusterId = svcClusterObj.Id
	return svcResp, nil
}
