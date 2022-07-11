package api

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/bbcoder-gh/api/router"
)

var onlyOnce sync.Once
var server *http.Server

func GetServer(port int) *http.Server {

	onlyOnce.Do(func() {

		//Configuring the web server
		server = &http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: router.Register(),
		}

	})

	return server
}
