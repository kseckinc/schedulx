package service

import (
	"github.com/galaxy-future/schedulx/api/types"
)

type CallBackNodeInitSvcReq struct {
	Instance *types.InstanceMeta `json:"instance"`
	Msg      string              `json:"msg"`
}

type TaskDescribeSvcReq struct {
	TaskId         int64                `json:"task_id"`
	InstanceStatus types.InstanceStatus `json:"instance_status"`
}

type TaskDescribeSvcResp struct {
	TaskDescribe *types.TaskDescribe `json:"task_describe"`
}

type TaskInstancesSvcReq struct {
	TaskId         int64                `json:"task_id"`
	InstanceStatus types.InstanceStatus `json:"instance_status"`
	PageNumber     int                  `json:"page_number"`
	PageSize       int                  `json:"page_size"`
}

type TaskInstancesSvcResp struct {
	InstancesList []*types.Instance `json:"instances_list"`
	Pager         *types.Pager      `json:"pager"`
}

type ExposeMountSvcReq struct {
	TaskId       int64                 `json:"task_id"`
	InstanceList []*types.InstanceInfo `json:"instance_list"`
	MountType    string                `json:"mount_type"`
	MountValue   string                `json:"mount_value"`
}

type ExposeMountSvcResp struct {
	TaskId       int64                 `json:"task_id"`
	InstanceList []*types.InstanceInfo `json:"instance_list"`
}

type ExposeUmountSvcReq struct {
	TaskId     int64  `json:"task_id"`
	Count      int64  `json:"count"`
	MountType  string `json:"mount_type"`
	MountValue string `json:"mount_value"`
}

type ExposeUmountSvcResp struct {
	TaskId       int64                 `json:"task_id"`
	InstanceList []*types.InstanceInfo `json:"instance_list"`
}
