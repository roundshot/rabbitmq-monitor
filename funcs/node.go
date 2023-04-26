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
	MemStats     `json:"memory"`
	Partitions   []string `json:"partitions"`
	Rawait       float64  `json:"io_read_avg_time"`
	Wawait       float64  `json:"io_write_avg_time"`
	Syncawait    float64  `json:"io_sync_avg_time"`
	MemUsed      int64    `json:"mem_used"`
	MemLimit     int64    `json:"mem_limit"`
	SocketsUsed  int64    `json:"sockets_used"`
	SocketsTotal int64    `json:"sockets_total"`
	FdUsed       int64    `json:"fd_used"`
	FdTotal      int64    `json:"fd_total"`
	ErlProcUsed  int64    `json:"proc_used