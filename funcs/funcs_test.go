package funcs

import (
	"testing"

	"github.com/barryz/rmqmonitor/g"
)

func TestGetExchanges(t *testing.T) {
	g.ParseConfig("../config.example.yml")
	//exchs, err := GetExchanges()
	//if err != nil {
	//