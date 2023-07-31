package system

import (
	"strings"
)

// Systemd is the systemd process control system.
type Systemd struct {
	name    string
	service string
}

// NewSystemd creates new Sy