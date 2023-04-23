package funcs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/barryz/rmqmonitor/g"
)

// MemStats ...
type MemStats struct {
	Total              int64 `json:"total"`
	ConnectionReaders  int64 `json:"connection_readers"`
	ConnectionWriters  int64 `json:"connection_writers"`
	ConnectionChannels int64 `json:"connection_channels"`
	ConnectionOther    int64 `json:"connection_other"`
	QueueProcs         int64 `json:"queue_procs"`
	QueueSlaveProcs    int64 `json:"queue_slave_procs"`
	Plugins            int64 `json:"plugins"`
	Mnesia             int64 `json:"mnesia"`
	MgmtDB             int64 `json:"mgmt_db"`
	MsgIndex           int64 `json:"msg_index"`
	Code               int64 `json:"code"`
	Atom               int64 `json:"atom"`
	Binary             int64 `json:"binary"`
}

// NodeStats ...
type NodeStats struct {
	MemStats     `json:"