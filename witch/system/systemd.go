package system

import (
	"strings"
)

// Systemd is the systemd process control system.
type Systemd struct {
	name    string
	service string
}

// NewSystemd creates new Systemd instance.
func NewSystemd(service string) *Systemd {
	return &Systemd{
		name:    "/bin/systemctl",
		service: