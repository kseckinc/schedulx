package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"gorm.io/gorm"

	"github.com/galaxy-future/schedulx/register/config/log"
	"github.com/galaxy-future/schedulx/register/constant"
	"github.com/galaxy-future/schedulx/repository/model/db"

	"github.com/galaxy-future/schedulx/api/types"
	jsoniter "github.com/json-iterator/go"

	"github.com/galaxy-future/schedulx/pkg/tool"
)

type ScheduleTemplateRepo struct {
}

type TemplateTaskLogic struct {
	TaskId         int64  `json:"task_id"`
	TmplExpandName string `json:"tmpl_expand_name"`
	ScheduleType   string `json:"schedule_type"`
	TaskInstCnt    int64  `json:"task_inst_cnt"`
	TaskExecType   string `json:"task_exec_type"`
	TaskExecOpr    string `json:"task_exec_opr"`
	BeginAt        string `json:"begin_at"`
	TimeCost       string `json:"time_cost"`
	TaskStatus     string `json:"task_status"`
	TaskStatusDesc string `json:"task_status_desc"`
	TaskStepDesc   string `json:"task_step_desc"`
}
type ExpandTemplateList struct {
	TmplExpandId    int64  `json:"tmpl_expand_id"`
	TmplExpandName  string `json:"tmpl_expand_name"`
	InstClusterName string `json:"inst_cluster_name"`
	IsContainer     bool   `json:"is_container"`
	RegisterType    string `json:"register_type"`
}

var scheduleTemplateRepoInst *ScheduleTemplateRepo
var scheduleTemplateRepoOnce sync.Once

func GetScheduleTemplateRepoInst() *ScheduleTemplateRepo {
	scheduleTemplateRepoOnce.Do(func() {
		scheduleTemplateRepoInst = &ScheduleTemplateRepo{}
	})

	return scheduleTemplateRepoInst
}

func (r *ScheduleTemplateRepo) GetSchedTmpl(schedTmplId int64) (*db.ScheduleTemplate, error) {
	var err error
	obj := &db.ScheduleTemplate{}
	if err = db.Get(schedTmplId, obj); err != nil {
		log.Logger.Error(err)
		return nil, err
	}

	return obj, nil
}

func (r *ScheduleTemplateRepo) GetSchedTmplBySvcClusterId(scId int64, schedType constant.ScheduleType) (*db.ScheduleTemplate, error) {
	var err error
	obj := &db.ScheduleTemplate{}
	where := map[string]interface{}{
		"service_cluster_id": scId,
		"schedule_type":      schedType,
	}
	if err = db.QueryFirst(where, obj); err != nil {
		log.Logger.Error(err)
		return nil, err
	}

	return obj, nil
}

func (r *ScheduleTemplateRepo) GetSchedReverseTmpl(schedTmplId int64) (*db.ScheduleTemplate, error) {
	var err error
	obj := &db.ScheduleTemplate{}
	if err = db.Get(schedTmplId, obj); err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	if obj.ReverseSchedTmplId == 0 {
		err = errors.New("reverse_sched_tmpl_id not found")
		log.Logger.Error(err)
		return nil, err
	}
	if err = db.Get(obj.ReverseSchedTmplId, obj); err != nil {
		log.Logger.Error(err)
		return nil, err
	}

	return obj, nil
}

func (r *ScheduleTemplateRepo) GetSchedReverseTmplBySvcClusterId(scId int64) (*db.ScheduleTemplate, error) {
	var err error
	obj := &db.ScheduleTemplate{}
	where := map[string]interface{}{
		"service_cluster_id": scId,
	}
	if err = db.QueryFirst(where, obj); err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	if obj.ReverseSchedTmplId == 0 {
		err = errors.New("reverse_sched_tmpl_id not found")
		log.Logger.Error(err)
		return nil, err
	}
	if err = db.Get(obj.ReverseSchedTmplId, obj); err != nil {
		log.Logger.Error(err)
		return nil, err
	}

	return obj, nil
}

func (r *ScheduleTemplateRepo) GetScheduleTempList(ctx context.Context, page, pageSize, schedule_cluster_id int) ([]TemplateTaskLogic, int64, error) {
	list := []TemplateTaskLogic{}

	scheduleTempModel := db.ScheduleTemplate{}
	where := map[string]interface{}{
		"service_cluster_id": schedule_cluster_id,
		"is_deleted":         0,
		"schedule_type":      constant.ScheduleTypeExpand,
	}
	err := db.QueryFirst(where, &scheduleTempModel)
	if err != nil {
		log.Logger.Errorf("db.Query table [schedule_template] error:%v", err)
	}
	scheduleTempMap := map[int64]string{
		scheduleTempModel.Id:                 string(scheduleTempModel.ScheduleType),
		scheduleTempModel.ReverseSchedTmplId: types.TaskShrink,
	}
	// 查询任务表信息
	taskModel := []db.Task{}
	taskWhere := map[string]interface{}{
		"sched_tmpl_id": []int64{scheduleTempModel.Id, scheduleTempModel.ReverseSchedTmplId},
	}
	fields := []string{"id", "task_step", "operator", "sched_tmpl_id", "task_status", "inst_cnt", "exec_type", "begin_at", "finish_at"}
	total, err := db.Query(taskWhere, page, pageSize, &taskModel, "id desc", fields, true)
	if err != nil {
		log.Logger.Errorf("table [task] queryAll error:%v", err)
	}
	// 构造关联task map
	list = make([]TemplateTaskLogic, len(taskModel))
	for index, item := range taskModel {
		var min, sec int64
		costTime := 0.0
		if item.FinishAt != nil {
			costTime = item.FinishAt.Sub(item.BeginAt).Seconds()
			min, sec = tool.SecondsToInt64(costTime)
		}
		list[index] = TemplateTaskLogic{
			TaskId:         item.Id,
			TmplExpandName: scheduleTempModel.TmplName,
			ScheduleType:   scheduleTempMap[item.SchedTmplId],
			TaskExecType:   item.ExecType,
			TaskStatus:     item.TaskStatus,
			TaskStatusDesc: types.TaskStatusDesc(item.TaskStatus),
			TaskStepDesc:   types.TaskStepDesc(item.TaskStep),
			TaskInstCnt:    item.InstCnt,
			TaskExecOpr:    item.Operator,
			BeginAt:        item.BeginAt.Format("2006-01-02 15:04:05"),
			TimeCost:       fmt.Sprintf("%d 分钟 %d 秒", min, sec),
		}
	}
	return list, total, nil
}

func (r *ScheduleTemplateRepo) GetExpandList(ctx context.Context, serviceName string, page, pageSize, serviceClusterId int) ([]ExpandTemplateList, int64, error) {
	var err error
	list := []ExpandTemplateList{}
	templateWhere := map[string]interface{}{
		"service_name":       serviceName,
		"service_cluster_id": serviceClusterId,
		"schedule_type":      constant.ScheduleTypeExpand,
		"is_deleted":         0,
	}
	templateModel := []db.ScheduleTemplate{}
	count, err := db.Query(templateWhere, page, pageSize, &templateModel, "id asc", nil, true)
	if err != nil {
		log.Logger.Errorf("db.Query table [schedule_tmeplate]error:%v", err)
	}
	if count == 0 {
		err = gorm.ErrRecordNotFound
		log.Logger.Error(err)
		return nil, 0, err
	}

	list = make([]ExpandTemplateList, len(templateModel))
	for index, item := range templateModel {
		instrIds := strings.Split(strings.Trim(item.InstrGroup, "[]"), ",")
		// 查询指令
		instrWhere := map[string]interface{}{
			"id": instrIds,
		}
		instrModel := []db.Instruction{}
		//查询指令步骤表
		err = db.QueryAll(instrWhere, &instrModel, "id asc", []string{"id", "params", "instr_action"})
		if err != nil {
			log.Logger.Errorf("db.Query table [instruction] error:%v", err)
			return nil, 0, err
		}
		IsContainer := false
		registerType := ""
		envBase := &types.BaseEnv{}
		mountALB := &types.ParamsMount{}
		for _, val := range instrModel {
			log.Logger.Infof("instrution id:%v,params info:%v", val.Id, val.Params)
			bytesParams := []byte(val.Params)
			if val.InstrAction == types.ENVInt && len(bytesParams) > 0 {
				err = jsoniter.Unmarshal(bytesParams, envBase)
				if err != nil {
					log.Logger.Errorf("table [instrution] nodeact.initbase error:%v", err)
					return nil, 0, err
				}
				if envBase.IsContainer {
					IsContainer = true
				}

			} else {
				err = errors.New("instruction.params [nodeact.initbase] invalid")
				log.Logger.Errorf("table [instrution] nodeact.initbase error:%v", err)
				//return nil, 0, err
			}

			if val.InstrAction == types.MountTypeSLB && len(bytesParams) > 0 {
				err = jsoniter.Unmarshal(bytesParams, mountALB)
				if err != nil {
					log.Logger.Errorf("table [instrution] mount.slb params:%v error:%v", val.Params, err)
					return nil, 0, err
				}

				if mountALB.MountType == strings.ToUpper(types.MountValueNginx) {
					registerType = strings.ToUpper(types.MountValueNginx)
				}
				if mountALB.MountType == strings.ToUpper(types.MountValueALB) {
					registerType = strings.ToUpper(types.MountValueALB)
				}
			} else {
				err = errors.New("instruction.params [slb] invalid")
				log.Logger.Errorf("table [instrution] val.InstrAction:%v mount.slb error:%v", val.InstrAction, err)
				continue
				//return nil, 0, err
			}
		}
		list[index] = ExpandTemplateList{
			TmplExpandName:  item.TmplName,
			TmplExpandId:    item.Id,
			InstClusterName: item.BridgxClusname,
			IsContainer:     IsContainer,
			RegisterType:    registerType,
		}
	}
	return list, count, nil
}

// 更新数据表
func (schtr *ScheduleTemplateRepo) Delete(ctx context.Context, tmplIds []int64) (int64, error) {
	templateModel := &db.ScheduleTemplate{}
	// 可以重复删除，不用事物加快效率
	templateListModel := []db.ScheduleTemplate{}
	err := db.Gets(tmplIds, &templateListModel)
	if err != nil && strings.Contains(err.Error(), "not record") {
		log.Logger.Errorf("error:%v", err)
	}

	for _, item := range templateListModel {
		tmplIds = append(tmplIds, item.ReverseSchedTmplId)
	}
	ret, err := db.Updates(templateModel, map[string]interface{}{"id": tmplIds}, map[string]interface{}{"is_deleted": 1}, nil)

	return ret, err
}
