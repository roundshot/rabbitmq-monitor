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
	Redeliver  QueueRate `json:"redeliver_details"`
}

// QueueMap ...
type QueueMap struct {
	Memory          int64       `json:"memory"`
	Messages        int64       `json:"messages"`
	MessagesReady   int64       `json:"messages_ready"`
	MessagesUnacked int64       `json:"messages_unacknowledged"`
	ConsumerUtil    interface{} `json:"consumer_utilisation"`
	Consumers       int64       `json:"consumers"`
	Status          string      `json:"state"`
	Name            string      `json:"name"`
	Vhost           string      `json:"vhost"`
	AutoDelete      bool        `json:"auto_delete"`
	QueueMsgStat    `json:"message_stats"`
}

func filterQueue(q *QueueMap) bool {
	isIgnore := false
	ignores := g.Config().Ignores

	for _, i := range ignores {
		if strings.Contains(strings.ToLower(q.Name), i) {
			isIgnore = true
		} else