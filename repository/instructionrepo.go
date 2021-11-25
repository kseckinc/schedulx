package repository

import (
	"context"
	"sync"

	"github.com/galaxy-future/schedulx/register/config"
	"github.com/galaxy-future/schedulx/register/config/log"
	"github.com/galaxy-future/schedulx/repository/model/db"
	"gorm.io/gorm"
)

type InstrRepo struct {
}

var instrRepoInst *InstrRepo
var instrRepoOnce sync.Once

func GetInstrRepoInst() *InstrRepo {
	instrRepoOnce.Do(func() {
		instrRepoInst = &InstrRepo{}
	})
	return instrRepoInst
}

func (r *InstrRepo) GetInstr(ctx context.Context, instrId int64) (*db.Instruction, error) {
	var err error
	obj := &db.Instruction{}
	err = db.Get(instrId, obj)
	if err != nil {
		log.Logger.Error(err.Error())
		return nil, err
	}
	return obj, nil
}

func (r *InstrRepo) DeleteByTmplExpandId(ctx context.Context, tmplExpandId int64, dbo *gorm.DB) error {
	var err error
	where := map[string]interface{}{
		"tmpl_id": tmplExpandId,
	}
	updates := map[string]interface{}{
		"is_deleted": 1,
	}
	rowsAffected, err := db.Updates(&db.Instruction{}, where, updates, dbo)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	if rowsAffected == 0 {
		err = config.ErrRowsAffectedInvalid
		log.Logger.Error(err)
		return err
	}

	return nil
}
