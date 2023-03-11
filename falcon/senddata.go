package falcon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/barryz/rmqmonitor/g"
)

func sendData(data []*MetaData) (resp []byte, err error) {
	debug := g.Config().Debug
	js, err := json.Marshal(data)
	if err != nil {
		return
	}

	if debug {
		log.Print