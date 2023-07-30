
package system

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

var (
	stopWaitSecs = 5
)

// System is the interface of process control system.
type System interface {
	// IsAlive checks process is alive.
	IsAlive() (int, bool)
	// Start starts process.
	Start() (bool, error)
	// Start restart process.
	Restart() (bool, error)
	// Stop stops process.
	Stop() bool
}

// Stats represents a stats interface
type Stats interface {
	// Reset the RabbitMQ StatsDB, A node may be randomly selected
	Reset() (bool, string, error)
	// Terminate the RabbitMQ StatsDB, Not to choose the node
	Terminate() (bool, string, error)
	// Crash the RabbitMQ StatsDB, Not to choose the node
	Crash() (bool, string, error)
}

// Action is the system action.
type Action struct {
	Name string `json:"name"`
}

// ActionStatus is the status of action.
type ActionStatus struct {
	Status bool   `json:"status"`
	Text   string `json:"text"`
}

// SysController controls the System.
type SysController struct {
	System
}
