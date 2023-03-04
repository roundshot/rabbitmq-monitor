package falcon

import (
	"testing"

	"github.com/barryz/rmqmonitor/g"
)

func TestGetStatsDB(t *testing.T) {
	g.ParseConfig("../config.example.yml")

	//ov, err := funcs.GetOve