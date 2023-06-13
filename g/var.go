package g

import (
	"github.com/barryz/rmqmonitor/utils"
)

// StatsDB stats management database
type StatsDB struct {
	CurrentLocate   string `json:"current_locate"`
	PreviousLocate  string `j