package service

import (
	"context"
	"sync"

	"github.com/galaxy-future/schedulx/pkg/tool"
	"github.com/galaxy-future/schedulx/register/config/log"
	"github.com/galaxy-future/schedulx/repository"
)

type NodeService struct {
}

var nodeServiceInstance *NodeService
var nodeOnce sync.Once

func GetNodeSvcInst() *NodeService {
	nodeOnce.Do(func() {
		nodeServiceInstance = &NodeService{}
	})
	return nodeServiceInstance
}

func (s *NodeService) entryLog(ctx context.Context, method string, req interface{}) {
	log.Logger.Infof("entry log | method[%s] | req:%s", method, tool.ToJson(req))
}

func (s *NodeService) exitLog(ctx context.Context, method string, req, resp interface{}, err error) {
	log.Logger.Infof("exit log | method[%s] | req:%s | resp:%s | err:%v", method, tool.ToJson(req), tool.ToJson(resp), err)
}

func (s *NodeService) UpdateNode(ctx context.Context, svcReq *CallBackNodeInitSvcReq) error {
	var err error
	s.entryLog(ctx, "UpdateNode", svcReq) // todo 日志脱敏
	defer func() {
		s.exitLog(ctx, "UpdateNode", svcReq, nil, err)
	}()
	repo := repository.GetInstanceRepoIns()
	if err = repo.UpdateInst(ctx, svcReq.Instance.TaskId, svcReq.Instance.InstanceId, svcReq.Instance.InstanceStatus, svcReq.Msg); err != nil {
		return err
	}

	return nil
}
