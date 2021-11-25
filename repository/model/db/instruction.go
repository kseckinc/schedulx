package db

import (
	"time"

	"github.com/galaxy-future/schedulx/api/types"
)

type Instruction struct {
	Id          int64        `gorm:"column:id" json:"id"`
	Cmd         string       `gorm:"column:cmd" json:"cmd"`
	Params      string       `gorm:"column:params" json:"params"`
	InstrAction types.Action `gorm:"column:instr_action" json:"instr_action"`
	TmplId      int64        `gorm:"column:tmpl_id" json:"tmpl_id"`
	IsDeleted   int8         `gorm:"column:is_deleted" json:"is_deleted"`
	CreateAt    *time.Time   `gorm:"column:create_at" json:"create_at"`
	UpdateAt    *time.Time   `gorm:"column:update_at" json:"update_at"`
}

func (t *Instruction) TableName() string {
	return "instruction"
}
