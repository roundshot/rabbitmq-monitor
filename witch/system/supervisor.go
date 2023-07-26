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
	return &Supervis