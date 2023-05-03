package funcs

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/barryz/rmqmonitor/g"
)

// QueueRate ...
type QueueRate struct {
	Rate float64 `json:"rate"`
}

// QueueMsgStat ...
type QueueMsgStat struct {
	Publish    QueueRate `json:"publis