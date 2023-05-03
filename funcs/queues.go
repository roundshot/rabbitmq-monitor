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
	Publish    QueueRate `json:"publish_details"`
	DeliverGet QueueRate `json:"deliver_get_details"`
	Ack        QueueRate `json:"ack_details"`
	Redeliver  Queue