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

// Server is the system R