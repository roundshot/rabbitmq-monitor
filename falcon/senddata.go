package falcon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/barryz/rmqmonitor/g"
)

func sendData(data []*MetaData) (resp []byte, er