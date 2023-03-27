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

// GetAlive ...
func GetAlive() (aliveness *Aliveness, err error) {
	service := "aliveness-test/%2f"

	res, err := g.RabbitAPI(service)
	if err != nil {
		err = fmt.Errorf("[ERROR]: get rabbitmq aliveness fail due to %s", err.Error())
		return
	}

	err = json.Unmarshal(res