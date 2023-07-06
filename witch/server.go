package witch

import (
	"log"
	"net/http"
	"strings"

	"github.com/barryz/rmqmonitor/g"
	"github.com/barryz/rmqmonitor/witch/system"

	"github.com/braintree/manners"
	"github.com/go-martini/martini"
	mauth "github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"
)

// Server is the system RESTFul web server.
type Server struct {
	addr string
	m    *martini.ClassicMartini
}

// NewServer inits a system RESTful web server.
func NewServer(addr string, sysControl *system.SysController, statsControl *system.StatsController, cfg *g.GlobalConfig) *Server {
	ser := &Server{
		addr: addr,
		m:    martini.Classic(),
	}
	authFu