package g

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

// GetHost get hostname (syscall)
func GetHost() string {
	return Config().Hostname
}

// GetAPIUrl 