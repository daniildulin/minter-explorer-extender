package models

import "time"

type Block struct {
	ID           uint        `gorm:"primary_key"`
	Height       uint        `json:"height"       gorm:"type:bigint;unique_index"`
	Timestamp    int64       `json:"timestamp"    gorm:"type:bigint"`
	TxCount      uint        `json:"tx_count"`
	Size         uint        `json:"size"`
	BlockTime    float64     `json:"block_time"   gorm:"type:numeric(20, 10)"`
	Hash         string      `json:"hash"         gorm:"type:varchar(255)"`
	BlockReward  string      `json:"block_reward" gorm:"type:numeric(300, 0)"`
	Validators   []Validator `gorm:"many2many:block_validator;"`
	Transactions []Transaction
	Rewards      []Reward
	Slashes      []Slash
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
