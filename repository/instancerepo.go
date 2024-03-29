package repository

import (
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/galaxy-future/schedulx/pkg/tool"

	"github.com/galaxy-future/schedulx/register/config"

	"github.com/galaxy-future/schedulx/api/types"
	"github.com/galaxy-future/schedulx/register/config/log"
	"github.com/galaxy-future/schedulx/repository/model/db"
)

type instanceRepo struct {
}

var (
	instanceRepoIns  *instanceRepo
	instanceRepoOnce sync.Once
)

func GetInstanceRepoIns() *instanceRepo {
	instanceRepoOnce.Do(
		func() {
			instanceRepoIns = &instanceRepo{}
		})
	return instanceRepoIns
}

func (r *instanceRepo) GetInsListCond(ctx context.Context, page, pageSize int, condition map[string]interface{}) ([]db.Instance, int64) {
	insList := []db.Instance{}
	total, err := db.Query(condition, page, pageSize, insList, "id asc", []string{}, true)
	if err != nil {
		log.Logger.Errorf("error:%v", err)
	}
	return insList, total
}

// UpdateStatus 更新实例状态
func (r *instanceRepo) UpdateStatus(ctx context.Context, status, taskId string, ipInner string) (int64, error) {
	instance := &db.Instance{}
	ret, err := db.Updates(instance, map[string]interface{}{"task_id": taskId, "ip_inner": ipInner}, map[string]interface{}{"instance_status": status}, nil)
	if err != nil {
		log.Logger.Errorf("db.Updates error:%v", err)
	}
	return ret, err
}

// BatchUpdateStatus 批量更新数据库表
func (r *instanceRepo) BatchUpdateStatus(ctx context.Context, ipInner []string, taskId int64, status types.InstanceStatus) (int64, error) {
	instance := &db.Instance{}
	where := map[string]interface{}{
		"task_id":  taskId,
		"ip_inner": ipInner,
	}
	columns := map[string]interface{}{
		"instance_status": status,
	}
	records, err := db.Updates(instance, where, columns, nil)
	if err != nil {
		log.Logger.Errorf("db.updates error %v:", err)
	}
	return records, err
}

// GetInsInfoWithIps 通过ip列表查询实例
func (r *instanceRepo) GetInsInfoWithIps(ctx context.Context, taskId string, ipList []string) []db.Instance {
	insList := &[]db.Instance{}
	where := map[string]interface{}{
		"ip_inner": ipList,
		"task_id":  taskId,
	}
	err := db.QueryAll(where, insList, "id asc", []string{"ip_inner", "instance_id"})
	if err != nil {
		log.Logger.Errorf(" GetInsInfo error:%v", err)
	}
	return *insList
}

func (r *instanceRepo) UpInsertInstanceBatch(ctx context.Context, instance []*types.InstanceInfo, taskId int64) error {
	var err error
	var instanceList []*db.Instance
	where := map[string]interface{}{
		"task_id": taskId,
	}
	if err = db.QueryAll(where, &instanceList, "id", nil); err != nil {
		log.Logger.Error("query all:", err)
		return err
	}
	existInstance := make(map[string]int, len(instanceList))
	for _, ins := range instanceList {
		existInstance[ins.InstanceId] = 1
	}
	// todo 是否对 existInstance 实例的 instance_status 做修改？
	var instanceBatch []*db.Instance
	for _, ins := range instance {
		if _, ok := existInstance[ins.InstanceId]; ok {
			continue
		}
		instance := &db.Instance{
			TaskId:         taskId,
			InstanceId:     ins.InstanceId,
			InstanceStatus: types.InstanceStatusInit,
			IpInner:        ins.IpInner,
			IpOuter:        ins.IpOuter,
			//CreateAt:       time.Now(),
		}
		instanceBatch = append(instanceBatch, instance)
	}
	if err = db.BatchCreate(instanceBatch, nil); err != nil {
		log.Logger.Error("batch create:", err)
		return err
	}

	return nil
}

func (r *instanceRepo) UpdateInst(ctx context.Context, taskId int64, instId string, instStatus types.InstanceStatus, msg string) error {
	var err error
	instance := &db.Instance{}
	where := map[string]interface{}{
		"task_id":     taskId,
		"instance_id": instId,
	}
	if err = db.QueryFirst(where, instance); err != nil {
		log.Logger.Error("query all:", err)
		return err
	}

	data := map[string]interface{}{
		"instance_status": instStatus,
		"msg":             tool.SubStr(msg, 100),
	}
	var rowsAffected int64
	if rowsAffected, err = db.Updates(&db.Instance{}, where, data, nil); err != nil {
		log.Logger.Error("db table [instance] updates:", err)
		return err
	}
	if rowsAffected != 1 {
		err = config.ErrRowsAffectedInvalid
		log.Logger.Error(err)
		return err
	}

	return nil
}

func (r *instanceRepo) InstsQueryByTaskId(ctx context.Context, taskId int64, instanceStatus types.InstanceStatus, fields []string) ([]*db.Instance, error) {
	var err error
	var instances []*db.Instance
	where := map[string]interface{}{
		"task_id": taskId,
	}
	if !strings.EqualFold(string(instanceStatus), "") {
		where["instance_status"] = instanceStatus
	}

	err = db.QueryAll(where, &instances, "id", fields)
	if err != nil {
		log.Logger.Error("db queryall:", err)
		return nil, err
	}

	return instances, err
}

func (r *instanceRepo) InstsQueryCntLimit(ctx context.Context, taskId int64, instanceStatus types.InstanceStatus, fields []string, limitCnt int) ([]*db.Instance, error) {
	var err error
	var instances []*db.Instance
	where := map[string]interface{}{
		"task_id": taskId,
	}
	if !strings.EqualFold(string(instanceStatus), "") {
		where["instance_status"] = instanceStatus
	}

	err = db.QueryLimit(where, &instances, "id", fields, limitCnt)
	if err != nil {
		log.Logger.Error("db queryall:", err)
		return nil, err
	}

	return instances, err
}

func (r *instanceRepo) QueryInstsToUmount(ctx context.Context, taskId int64, instanceStatus types.InstanceStatus, limit int) (int64, []*types.InstanceInfo, error) {
	var err error
	fields := []string{
		"ip_inner",
		"ip_outer",
		"instance_id",
	}
	insts, err := r.InstsQueryCntLimit(ctx, taskId, instanceStatus, fields, limit)
	if err != nil {
		return 0, nil, err
	}
	if len(insts) == 0 {
		err = errors.New("no record found")
		log.Logger.Error(err)
		return 0, nil, err
	}
	instList := make([]*types.InstanceInfo, 0, len(insts))
	for _, i := range insts {
		inst := &types.InstanceInfo{
			IpInner:    i.IpInner,
			IpOuter:    i.IpOuter,
			InstanceId: i.InstanceId,
		}
		instList = append(instList, inst)
	}

	return taskId, instList, nil
}

func (r *instanceRepo) InstsQueryByPage(ctx context.Context, taskId int64, instanceStatus types.InstanceStatus, pageSize, PageNumber int, fields []string) ([]*db.Instance, int64, error) {
	var err error
	var instances []*db.Instance
	where := map[string]interface{}{
		"task_id": taskId,
	}
	if !strings.EqualFold(string(instanceStatus), "") {
		where["instance_status"] = instanceStatus
	}

	cnt, err := db.Query(where, PageNumber, pageSize, &instances, "id", fields, true)
	if err != nil {
		log.Logger.Error("db queryall:", err)
		return nil, 0, err
	}

	return instances, cnt, err
}
