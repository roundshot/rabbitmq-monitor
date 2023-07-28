package system

import (
	"strings"
)

// Supervisor is the supervisor process control system.
type Supervisor struct {
	name    string
	service string
}

// NewSupervisor creates new Supervisor instance.
func NewSupervisor(service string) *Supervisor {
	return &Supervisor{
		name:    "/usr/bin/supervisorctl",
		service: service,
	}
}

// IsAlive gets results from `supervisorctl status [service]`
func (s *Supervisor) IsAlive() (i