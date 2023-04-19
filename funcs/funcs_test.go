package funcs

import (
	"testing"

	"github.com/barryz/rmqmonitor/g"
)

func TestGetExchanges(t *testing.T) {
	g.ParseConfig("../config.examp