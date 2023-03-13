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
		log.Printf("agent api received %d metrics", len(data))
	}

	res, err := http.Post(g.Config().Falcon.API, "Content-Type: application/json", bytes.NewBuffer(js))
	if err != nil {
		err = fmt.Errorf("[ERROR]: sent data to falcon agent api fail due to %s", err.Error())
		return
	}

	defer res.Body.Close()