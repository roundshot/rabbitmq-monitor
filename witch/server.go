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
	authFunc := mauth.BasicFunc(func(username, password string) bool {
		pwd, ok := cfg.Witch.Auth[username]
		return ok && pwd == password
	}).(func(http.ResponseWriter, *http.Request, martini.Context))
	ser.m.Map(sysControl)
	ser.m.Map(statsControl)
	ser.m.Use(authInclusive("/api", authFunc))
	ser.m.Use(render.Renderer(render.Options{}))
	// start|stop|restart RabbitMQ process(other process)  via supervisor or systemd
	ser.m.Put("/api/app/actions", sysAction)
	// forced to stop RabbitMQ process(other process) via sent SIGTERM syscall signal
	ser.m.Get("/api/app/fstop", procForceStop)
	// get current RabbitMQ statistic db node location
	ser.m.Get("/api/stats", statsInfo)
	// reset|crash|terminate current RabbitMQ statistic db node
	ser.m.Put("/api/stats/actions", statsAction)
	return ser
}

// Start starts the server.