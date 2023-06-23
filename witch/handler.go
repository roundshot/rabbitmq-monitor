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
	ErrServerError = errors.New("Internal Server Error")
	// ErrBadRequest is bad request error.
	ErrBadRequest = errors.New("Bad Request")
)

func sysAction(control *system.SysController, req *http.Request, r render.Render) {
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("[ERROR] Read request body error: %s", err)
		r.JSON(http.StatusInternalServe