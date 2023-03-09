
package falcon

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/barryz/rmqmonitor/funcs"
	"github.com/barryz/rmqmonitor/g"
)

var (
	statsDB = g.NewStatsDB()
)

const (
	overviewPrefix = "rabbitmq.overview."
	queuePrefix    = "rabbitmq.queue."
	exchangePrefix = "rabbitmq.exchange."
)

// MetaData meta data
type MetaData struct {
	Endpoint    string      `json:"endpoint"`
	Metric      string      `json:"metric"`
	Value       interface{} `json:"value"`
	CounterType string      `json:"counterType"`
	Tags        string      `json:"tags"`
	Timestamp   int64       `json:"timestamp"`
	Step        int64       `json:"step"`
}

// NewMetric create an new metric
func NewMetric(name string, value interface{}, tags string) *MetaData {
	host := g.GetHost()
	return &MetaData{
		Metric:      name,
		Endpoint:    host,
		CounterType: fmt.Sprintf("GAUGE"),
		Tags:        tags,
		Timestamp:   time.Now().Unix(),