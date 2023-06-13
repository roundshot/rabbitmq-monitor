package g

import (
	"github.com/barryz/rmqmonitor/utils"
)

// StatsDB stats management database
type StatsDB struct {
	CurrentLocate   string `json:"current_locate"`
	PreviousLocate  string `json:"previous_locate"`
	LastChangeTime  string `json:"last_change"`
	LastCollectTime string `json:"last_collect"`
}

// NewStatsDB