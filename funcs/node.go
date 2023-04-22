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
	QueueSlaveProcs    i