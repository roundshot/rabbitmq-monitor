package funcs

import (
	"encoding/json"
	"fmt"

	"github.com/barryz/rmqmonitor/g"
)

// Aliveness ...
type Aliveness struct {
	Status string `json:"status"`
}

// G