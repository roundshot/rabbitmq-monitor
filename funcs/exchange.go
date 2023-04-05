package funcs

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/barryz/rmqmonitor/g"
)

// VHostName name of vHost
type VHostName struct {
	Name string `json:"name"`
}

// ExchangeInfo information for exchange
type ExchangeInfo struct {
	Name     string `json:"name"`
	VHost    string `json:"vhost"`
	MsgStats struct {
		Confirm    int64 `json:"confirm"`
		PublishIn  int64 `js