package witch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/barryz/rmqmonitor/falcon"
	"github.com/barryz/rmqmonitor/g"
	"github.com/barryz/rmqmonitor/witch/system"

	"github.com/martini-contrib/render"
)

var (
	// ErrServerError is internal server error.
	ErrServerError = err