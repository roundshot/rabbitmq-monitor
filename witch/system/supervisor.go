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
func (s *Supervisor) IsAlive() (int, bool) {
	r, err := ExecCommand(s.name, []string{"status", s.service})
	if err != nil {
		return -1, false
	}
	return -1, strings.Contains(r, "RUNNING")
}

// Start executes `supervisorctl start [service]`
func (s *Supervisor) Start() (bool, error) {
	_, err := ExecCommand(s.name, []string{"star